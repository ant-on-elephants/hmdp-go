package controller

import (
	"github.com/gin-gonic/gin"
	"hmdp-go/common/codes"
	"hmdp-go/common/logger"
	"hmdp-go/service"
	"net/http"
)

// ShopType 注入IShopTypeService
type ShopType struct {
	Log     logger.ILogger           `inject:""`
	Service service.IShopTypeService `inject:""`
}

func (a *ShopType) QueryTypeList(c *gin.Context) {

	// 根据类型分页查询
	list, err := a.Service.GetShopTypeList("sort asc")
	if err != nil {
		RespCommonFail(c)
		return
	}

	// 返回数据
	RespData(c, http.StatusOK, codes.SUCCESS, list)
}
