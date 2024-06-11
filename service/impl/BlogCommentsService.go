package impl

import (
	"hmdp-go/common/logger"
)

// BlogCommentsService 注入IDb
type BlogCommentsService struct {
	Log logger.ILogger `inject:""`
}
