package repository

import "hmdp-go/models"

// IShopRepository Shop接口定义
type IShopRepository interface {
	//GetTables 分页返回Shops
	GetTables(PageNum, PageSize int64, where interface{}) []*models.Shop
	//GetShop 根据id获取Shop
	GetShop(where interface{}) *models.Shop
	//AddShop 新增Shop
	AddShop(shop *models.Shop) bool
	//GetShops 获取Shop
	GetShops(PageNum int64, PageSize int64, total *uint64, where interface{}) []*models.Shop

	Save(shop *models.Shop) (*models.Shop, error)

	//UpdateShopByWhere 更新Shop
	UpdateShopByWhere(where interface{}, update interface{}) bool
}
