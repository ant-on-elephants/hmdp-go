package impl

import (
	"hmdp-go/common/logger"
)

// FollowService 注入IDb
type FollowService struct {
	Log logger.ILogger `inject:""`
}
