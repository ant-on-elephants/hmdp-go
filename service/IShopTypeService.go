package service

import "hmdp-go/models"

// IShopTypeService ShopType接口定义
type IShopTypeService interface {
	//GetShopTypeList 获取ShopTypes
	GetShopTypeList(order interface{}) ([]*models.ShopType, error)
}
