package impl

import (
	"hmdp-go/common/logger"
)

// VoucherService 注入IDb
type VoucherService struct {
	Log logger.ILogger `inject:""`
}
