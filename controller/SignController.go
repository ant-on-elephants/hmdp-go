package controller

import (
	"hmdp-go/common/logger"
	"hmdp-go/service"
)

// Sign 注入ISignService
type Sign struct {
	Log     logger.ILogger       
	Service service.ISignService 
}
