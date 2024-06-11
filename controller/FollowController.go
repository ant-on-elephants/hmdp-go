package controller

import (
	"hmdp-go/common/logger"
	"hmdp-go/service"
)

// Follow 注入IFollowService
type Follow struct {
	Log     logger.ILogger         
	Service service.IFollowService 
}
