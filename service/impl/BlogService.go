package impl

import (
	"hmdp-go/common/logger"
)

// BlogService 注入IDb
type BlogService struct {
	Log logger.ILogger `inject:""`
}
