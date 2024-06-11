package jwt

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"hmdp-go/repository"
	"log"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"

	"hmdp-go/common/codes"
	"hmdp-go/common/setting"
	"hmdp-go/models"
)

// app 程序配置
var app = setting.Config.APP

const CtxCurrentUser = "CURRENT_USER"

type CurrentUser struct {
	Id       int64  // 用户id
	NickName string // 名称
	Icon     string // 头像
}

// JWT 注入IService
type JWT struct {
	UserRepo repository.IUserRepository `inject:""`
}

// GinJWTMiddlewareInit 初始化中间件
func (j *JWT) GinJWTMiddlewareInit(jwtAuthorizator IAuthorizator) (authMiddleware *jwt.GinJWTMiddleware) {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour * 24,
		MaxRefresh:  time.Hour,
		IdentityKey: app.IdentityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			// 鉴权通过，将用户数据写入上下文
			if v, ok := data.(models.User); ok {
				return jwt.MapClaims{
					"id":        v.Id,
					"nick_name": v.NickName,
					"icon":      v.Icon,
				}
			}

			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} { // 解析并设置用户身份信息
			claims := jwt.ExtractClaims(c)
			//extracts identity from claims
			id := int64(claims["id"].(float64))
			nickname := claims["nick_name"].(string)
			icon := claims["icon"].(string)

			user := &CurrentUser{
				Id:       id,
				NickName: nickname,
				Icon:     icon,
			}
			return user
		},

		/**
		 * 登录功能
		 * @param loginForm 登录参数，包含手机号、验证码；或者手机号、密码
		 */
		Authenticator: func(c *gin.Context) (interface{}, error) {
			// TODO 实现登录功能

			data := make(map[string]interface{}) // 前端传入的参数
			err := c.ShouldBindBodyWith(&data, binding.JSON)
			if err != nil {
				return nil, codes.NewErrCode(codes.InvalidParams)
			}

			// todo
			phone := data["phone"].(string) // phone
			fmt.Println(phone)

			// 查询是否有此用户
			userInfo := j.UserRepo.GetUserByWhere("phone = " + phone)
			if userInfo.Id <= 0 {
				return nil, codes.NewErrCode(codes.ErrNotFoundUser)
			}

			user := CurrentUser{
				Id:       userInfo.Id,
				NickName: userInfo.NickName,
				Icon:     userInfo.Icon,
			}
			newCtx := context.WithValue(c.Request.Context(), CtxCurrentUser, user)
			c.Request = c.Request.WithContext(newCtx)
			return user, nil
		},
		//receives identity and handles authorization logic
		Authorizator: jwtAuthorizator.HandleAuthorizator,
		//handles unauthorized logic
		Unauthorized: func(c *gin.Context, code int, message string) { // 处理不进行授权的逻辑(未验证成功)
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// loginHandler处理成功后返回的这个
		LoginResponse: func(c *gin.Context, code int, message string, expireTime time.Time) {
			data := make(map[string]interface{})
			data["token"] = message
			data["expireTime"] = expireTime
			c.JSON(code, gin.H{
				"code": code,
				"data": data,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
	return
}

// NoRouteHandler 404 handler
func NoRouteHandler(c *gin.Context) {
	code := codes.PageNotFound
	c.JSON(404, gin.H{"code": code, "message": codes.GetMsg(code)})
}

// GetUser 获取登录用户信息
func GetUser(ctx context.Context) *CurrentUser {
	return ctx.Value(CtxCurrentUser).(*CurrentUser)
}
