package controller

import (
	"github.com/Masterminds/squirrel"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"hmdp-go/common/codes"
	"hmdp-go/common/constants"
	"hmdp-go/common/logger"
	"hmdp-go/common/middleware/jwt"
	"hmdp-go/repository"
	"hmdp-go/service"
	"net/http"
	"strconv"
)

// Blog 注入IBlogService
type Blog struct {
	Log      logger.ILogger
	Service  service.IBlogService
	BlogRepo repository.IBlogRepository
	UserRepo repository.IUserRepository
}

func (a *Blog) SaveBlog(c *gin.Context) {
	data := make(map[string]interface{}) // 前端传入的参数
	err := c.ShouldBindBodyWith(&data, binding.JSON)
	if err != nil {
		RespUnsuccessful(c, http.StatusOK, codes.InvalidParams)
		return
	}

	blogId := data["blog_id"].(string) // blogId
	blog := a.BlogRepo.GetBlog("id = " + blogId)
	if blog.Id <= 0 {
		RespFail(c, http.StatusOK, codes.InvalidParams, "Blog未找到")
	}

	// 获取登录用户
	user := jwt.GetUser(c.Request.Context())

	blog.UserId = user.Id
	// 保存探店博文
	a.BlogRepo.UpdateBlog(blog)
	// 返回id
	RespData(c, http.StatusOK, codes.SUCCESS, blog.Id)
}

func (a *Blog) LikeBlog(c *gin.Context) {
	data := make(map[string]interface{}) // 前端传入的参数
	err := c.ShouldBindBodyWith(&data, binding.JSON)
	if err != nil {
		RespUnsuccessful(c, http.StatusOK, codes.InvalidParams)
		return
	}

	blogId := data["id"].(string) // blogId
	result := a.BlogRepo.UpdateBlogByWhere("id = "+blogId, "liked = liked + 1")
	if !result {
		RespFail(c, http.StatusOK, codes.InvalidParams, "Blog点赞失败")
		return
	}

	RespOk(c, http.StatusOK, codes.SUCCESS)
}

func (a *Blog) QueryMyBlog(c *gin.Context) {
	// 获取登录用户
	user := jwt.GetUser(c.Request.Context())

	value := c.Query("value")
	current, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		RespFail(c, http.StatusOK, codes.InvalidParams, err.Error())
		return
	}

	// 根据用户查询
	list := a.BlogRepo.GetTables(current, constants.MaxPageSize, squirrel.Eq{"user_id": user.Id})

	// 获取当前页数据
	RespData(c, http.StatusOK, codes.SUCCESS, list)
}

func (a *Blog) QueryHotBlog(c *gin.Context) {
	value := c.Query("value")
	current, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		RespFail(c, http.StatusOK, codes.InvalidParams, err.Error())
		return
	}

	query := squirrel.Select("*").OrderBy("liked -1") // 降序排序
	// 获取当前页数据
	list := a.BlogRepo.GetTables(current, constants.MaxPageSize, query)
	for _, blog := range list {
		user := a.UserRepo.GetUserByID(blog.UserId)
		blog.Name = user.NickName
		blog.Icon = user.Icon
	}

	RespData(c, http.StatusOK, codes.SUCCESS, list)
}
