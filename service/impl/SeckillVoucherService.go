package impl

import (
	"hmdp-go/common/logger"
)

// SeckillVoucherService 注入IDb
type SeckillVoucherService struct {
	Log logger.ILogger `inject:""`
}
