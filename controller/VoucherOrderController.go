package controller

import (
	"github.com/gin-gonic/gin"
	"hmdp-go/common/codes"
	"hmdp-go/common/logger"
	"hmdp-go/service"
	"net/http"
)

// VoucherOrder 注入IVoucherOrderService
type VoucherOrder struct {
	Log     logger.ILogger               
	Service service.IVoucherOrderService 
}

func (a *VoucherOrder) SeckillVoucher(c *gin.Context) {
	RespFail(c, http.StatusOK, codes.ERROR, "功能未完成")
}
