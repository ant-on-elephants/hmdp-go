package controller

import (
	"github.com/gin-gonic/gin"
	"hmdp-go/common/codes"
	"hmdp-go/common/logger"
	"hmdp-go/repository"
	"hmdp-go/service"
	"net/http"
	"time"
)

// User 注入IUserService
type User struct {
	Log          logger.ILogger                 
	Service      service.IUserService           
	UserInfoRepo repository.IUserInfoRepository 
}

// SendCode 发送手机验证码
func (a *User) SendCode(c *gin.Context) {
	// TODO 发送短信验证码并保存验证码
	RespFail(c, http.StatusOK, codes.ERROR, "功能未完成")
}

// Logout 退出登录
func (a *User) Logout(c *gin.Context) {
	// TODO 实现登出功能
	RespFail(c, http.StatusOK, codes.ERROR, "功能未完成")
}

// Me 获取当前登录的用户并返回
func (a *User) Me(c *gin.Context) {
	// TODO 获取当前登录的用户并返回
	RespFail(c, http.StatusOK, codes.ERROR, "功能未完成")
}

func (a *User) Info(c *gin.Context) {
	userInfo := a.UserInfoRepo.GetUserInfo("id = " + c.Query("userId"))
	if userInfo.UserId <= 0 {
		// 没有详情，应该是第一次查看详情
		RespData(c, http.StatusOK, codes.SUCCESS, nil)
		return
	}
	userInfo.CreateTime = time.Time{}
	userInfo.UpdateTime = time.Time{}
	// 返回
	RespData(c, http.StatusOK, codes.SUCCESS, userInfo)
}
