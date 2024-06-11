package controller

import (
	"hmdp-go/common/logger"
	"hmdp-go/service"
)

// BlogComments 注入IBlogCommentsService
type BlogComments struct {
	Log     logger.ILogger
	Service service.IBlogCommentsService
}
