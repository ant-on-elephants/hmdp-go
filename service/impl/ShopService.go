package impl

import (
	"hmdp-go/common/logger"
)

// ShopService 注入IDb
type ShopService struct {
	Log logger.ILogger `inject:""`
}
