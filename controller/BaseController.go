package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"hmdp-go/common/codes"
	"hmdp-go/common/setting"
)

// ResponseData 数据返回结构体
type ResponseData struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// ResponseSuccess 返回成功结构体
type ResponseSuccess struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// ResponseFail 返回成功结构体
type ResponseFail struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Detail string `json:"detail"`
}

// RespData 数据返回
func RespData(c *gin.Context, httpCode, code int, data interface{}) {
	resp := ResponseData{
		Code: code,
		Msg:  codes.GetMsg(code),
		Data: data,
	}
	RespJSON(c, httpCode, resp)
}

// RespOk 返回操作成功
func RespOk(c *gin.Context, httpCode, code int) {
	resp := ResponseSuccess{
		Code: code,
		Msg:  codes.GetMsg(code),
	}
	RespJSON(c, httpCode, resp)
}

// RespSuccess 返回操作成功
func RespSuccess(c *gin.Context) {
	resp := ResponseSuccess{
		Code: codes.SUCCESS,
		Msg:  codes.GetMsg(codes.SUCCESS),
	}
	RespJSON(c, http.StatusOK, resp)
}

// RespFail 返回操作失败
func RespFail(c *gin.Context, httpCode, code int, detail string) {
	resp := ResponseFail{
		Code:   code,
		Msg:    codes.GetMsg(code),
		Detail: detail,
	}
	RespJSON(c, httpCode, resp)
}

// RespCommonFail 返回操作失败
func RespCommonFail(c *gin.Context) {
	resp := ResponseFail{
		Code: codes.ERROR,
		Msg:  codes.GetMsg(codes.ERROR),
	}
	RespJSON(c, http.StatusInternalServerError, resp)
}

// RespUnsuccessful 返回操作失败
func RespUnsuccessful(c *gin.Context, httpCode, code int) {
	resp := ResponseFail{
		Code: code,
		Msg:  codes.GetMsg(code),
	}
	RespJSON(c, httpCode, resp)
}

// RespJSON 返回JSON数据
func RespJSON(c *gin.Context, httpCode int, resp interface{}) {
	c.JSON(httpCode, resp)
	c.Abort()
}

// GetPage 获取每页数量
func GetPage(c *gin.Context) (page, pagesize int) {
	page, _ = strconv.Atoi(c.Query("page"))
	pagesize, _ = strconv.Atoi(c.Query("limit"))
	if pagesize == 0 {
		pagesize = setting.Config.APP.Pagesize
	}
	if page == 0 {
		page = 1
	}
	return
}
