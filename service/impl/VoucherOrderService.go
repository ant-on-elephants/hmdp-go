package impl

import (
	"hmdp-go/common/logger"
)

// VoucherOrderService 注入IDb
type VoucherOrderService struct {
	Log logger.ILogger `inject:""`
}
