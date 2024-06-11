package repository

import "hmdp-go/models"

// IShopTypeRepository ShopType接口定义
type IShopTypeRepository interface {
	//GetTables 分页返回ShopTypes
	GetTables(PageNum, PageSize int64, where interface{}) []*models.ShopType
	//GetShopType 根据id获取ShopType
	GetShopType(where interface{}) *models.ShopType
	//AddShopType 新增ShopType
	AddShopType(shopType *models.ShopType) bool
	//GetShopTypes 获取ShopType
	GetShopTypes(PageNum int64, PageSize int64, total *uint64, where interface{}) []*models.ShopType
	//GetShopTypeList 获取ShopTypes
	GetShopTypeList(where interface{}, order interface{}) ([]*models.ShopType, error)
}
