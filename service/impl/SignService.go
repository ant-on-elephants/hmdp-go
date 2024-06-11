package impl

import (
	"hmdp-go/common/logger"
)

// SignService 注入IDb
type SignService struct {
	Log logger.ILogger `inject:""`
}
