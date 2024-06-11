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

// Shop 注入IShopService
type Shop struct {
	Log      logger.ILogger
	Service  service.IShopService
	ShopRepo repository.IShopRepository
}

// QueryShopById
/**
 * 根据id查询商铺信息
 * @param id 商铺id
 * @return 商铺详情数据
 */
func (a *Shop) QueryShopById(c *gin.Context) {
	value := c.Query("id")
	id, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		RespFail(c, http.StatusOK, codes.InvalidParams, err.Error())
		return
	}

	shop := a.ShopRepo.GetShop(squirrel.Eq{"id": id})
	if shop.Id <= 0 {
		RespFail(c, http.StatusOK, codes.InvalidParams, "Shop未找到")
		return
	}

	// 返回
	RespData(c, http.StatusOK, codes.SUCCESS, shop)
}

// SaveShop
/**
 * 新增商铺信息
 * @param shop 商铺数据
 * @return 商铺id
 */
func (a *Shop) SaveShop(c *gin.Context) {
	data := make(map[string]interface{}) // 前端传入的参数
	err := c.ShouldBindBodyWith(&data, binding.JSON)
	if err != nil {
		RespUnsuccessful(c, http.StatusOK, codes.InvalidParams)
		return
	}

	shop := data["shop"].(models.Shop)

	newShop, saveErr := a.ShopRepo.Save(&shop)
	if saveErr != nil {
		RespFail(c, http.StatusOK, codes.InvalidParams, saveErr.Error())
		return
	}

	// 返回
	RespData(c, http.StatusOK, codes.SUCCESS, newShop.Id)
}

// UpdateShop
/**
 * 更新商铺信息
 * @param shop 商铺数据
 * @return 无
 */
func (a *Shop) UpdateShop(c *gin.Context) {
	data := make(map[string]interface{}) // 前端传入的参数
	err := c.ShouldBindBodyWith(&data, binding.JSON)
	if err != nil {
		RespUnsuccessful(c, http.StatusOK, codes.InvalidParams)
		return
	}

	shop := data["shop"].(models.Shop)

	result := a.ShopRepo.UpdateShopByWhere(squirrel.Eq{"id": shop.Id}, shop)
	if !result {
		RespFail(c, http.StatusOK, codes.InvalidParams, "Shop更新失败")
		return
	}

	RespOk(c, http.StatusOK, codes.SUCCESS)
}

// QueryShopByType
/**
 * 根据商铺类型分页查询商铺信息
 * @param typeId 商铺类型
 * @param current 页码
 * @return 商铺列表
 */
func (a *Shop) QueryShopByType(c *gin.Context) {
	value := c.Query("typeId")
	typeId, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		RespFail(c, http.StatusOK, codes.InvalidParams, err.Error())
		return
	}

	value = c.Query("current")
	current, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		RespFail(c, http.StatusOK, codes.InvalidParams, err.Error())
		return
	}

	// 根据类型分页查询
	list := a.ShopRepo.GetTables(current, constants.MaxPageSize, squirrel.Eq{"type_id": typeId})

	// 返回数据
	RespData(c, http.StatusOK, codes.SUCCESS, list)
}

// QueryShopByName
/**
 * 根据商铺名称关键字分页查询商铺信息
 * @param name 商铺名称关键字
 * @param current 页码
 * @return 商铺列表
 */
func (a *Shop) QueryShopByName(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		RespUnsuccessful(c, http.StatusOK, codes.InvalidParams)
		return
	}

	value := c.Query("current")
	current, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		RespFail(c, http.StatusOK, codes.InvalidParams, err.Error())
		return
	}

	// 根据类型分页查询
	list := a.ShopRepo.GetTables(current, constants.MaxPageSize, squirrel.Like{"name": name})

	// 返回数据
	RespData(c, http.StatusOK, codes.SUCCESS, list)
}
