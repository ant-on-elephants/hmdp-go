package impl

import (
	"hmdp-go/common/logger"
)

// UserInfoService 注入IDb
type UserInfoService struct {
	Log logger.ILogger `inject:""`
}
