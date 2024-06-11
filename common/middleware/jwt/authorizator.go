package jwt

import (
	"github.com/gin-gonic/gin"
)

// IAuthorizator 授权规则接口
type IAuthorizator interface {
	HandleAuthorizator(data interface{}, c *gin.Context) bool
}

// AdminAuthorizator 管理员授权规则
type AdminAuthorizator struct {
}

// HandleAuthorizator 处理管理员授权规则
func (*AdminAuthorizator) HandleAuthorizator(data interface{}, c *gin.Context) bool {
	return true
}

// TestAuthorizator 测试用户授权规则
type TestAuthorizator struct {
}

// HandleAuthorizator 处理测试用户授权规则
func (*TestAuthorizator) HandleAuthorizator(data interface{}, c *gin.Context) bool {
	return true
}

// AllUserAuthorizator 普通用户授权规则
type AllUserAuthorizator struct {
}

// HandleAuthorizator 处理普通用户授权规则
func (*AllUserAuthorizator) HandleAuthorizator(data interface{}, c *gin.Context) bool {
	return true
}
