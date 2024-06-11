package impl

import (
	"hmdp-go/common/logger"
)

// UserService 注入IDb
type UserService struct {
	Log logger.ILogger `inject:""`
}
