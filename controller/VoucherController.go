package controller

import (
	"github.com/Masterminds/squirrel"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"hmdp-go/common/codes"
	"hmdp-go/common/constants"
	"hmdp-go/common/logger"
	"hmdp-go/models"
	"hmdp-go/repository"
	"hmdp-go/service"
	"net/http"
	"strconv"
)

// Voucher 注入IVoucherService
type Voucher struct {
	Log         logger.ILogger
	Service     service.IVoucherService
	VoucherRepo repository.IVoucherRepository
}

func (a *Voucher) AddVoucher(c *gin.Context) {
	data := make(map[string]interface{}) // 前端传入的参数
	err := c.ShouldBindBodyWith(&data, binding.JSON)
	if err != nil {
		RespUnsuccessful(c, http.StatusOK, codes.InvalidParams)
		return
	}

	voucher := data["voucher"].(models.Voucher)

	newVoucher, saveErr := a.VoucherRepo.Save(&voucher)
	if saveErr != nil {
		RespFail(c, http.StatusOK, codes.InvalidParams, saveErr.Error())
		return
	}

	RespData(c, http.StatusOK, codes.SUCCESS, newVoucher.Id)
}

func (a *Voucher) AddSeckillVoucher(c *gin.Context) {
	data := make(map[string]interface{}) // 前端传入的参数
	err := c.ShouldBindBodyWith(&data, binding.JSON)
	if err != nil {
		RespUnsuccessful(c, http.StatusOK, codes.InvalidParams)
		return
	}

	voucher := data["voucher"].(models.Voucher)

	newVoucher, saveErr := a.VoucherRepo.Save(&voucher)
	if saveErr != nil {
		RespFail(c, http.StatusOK, codes.InvalidParams, saveErr.Error())
		return
	}

	RespData(c, http.StatusOK, codes.SUCCESS, newVoucher.Id)
}

func (a *Voucher) QueryVoucherOfShop(c *gin.Context) {
	shopId := c.Query("shopId")
	id, err := strconv.ParseInt(shopId, 10, 64)
	if err != nil {
		RespFail(c, http.StatusOK, codes.InvalidParams, err.Error())
		return
	}

	vouchers := a.VoucherRepo.GetTables(id, constants.MaxPageSize, squirrel.Like{"shopId": id})

	RespData(c, http.StatusOK, codes.SUCCESS, vouchers)
}
