package repository_impl

import (
	"hmdp-go/common/logger"
	"hmdp-go/models"
)

// ShopRepository 注入IDb
type ShopRepository struct {
	Log  logger.ILogger `inject:""`
	Base BaseRepository `inject:"inline"`
}

// GetTables 分页返回Shops
func (a *ShopRepository) GetTables(PageNum, PageSize int64, where interface{}) []*models.Shop {
	var shops []*models.Shop
	var total uint64
	err := a.Base.GetPages(&models.Shop{}, &shops, PageNum, PageSize, &total, "")
	if err != nil {
		a.Log.Errorf("GetTables函数出错：", err)
	}
	return shops
}

// GetShop 根据id获取Shop
func (a *ShopRepository) GetShop(where interface{}) *models.Shop {
	var shop models.Shop
	if err := a.Base.First(where, &shop); err != nil {
		a.Log.Errorf("未找到相关Shop", err)
	}
	return &shop
}

// AddShop 新增Shop
func (a *ShopRepository) AddShop(shop *models.Shop) bool {
	if err := a.Base.Save(shop); err != nil {
		a.Log.Errorf("添加Shop失败", err)
	}
	return true
}

// GetShops 获取Shop
func (a *ShopRepository) GetShops(PageNum int64, PageSize int64, total *uint64, where interface{}) []*models.Shop {
	var shops []*models.Shop
	err := a.Base.GetPages(&models.Shop{}, &shops, PageNum, PageSize, total, where, "ID desc")
	if err != nil {
		a.Log.Errorf("获取Shop信息失败", err)
	}
	return shops
}

// Save 保存
func (a *ShopRepository) Save(shop *models.Shop) (*models.Shop, error) {
	ret := a.Base.Source.DB().Save(shop)
	if ret.Error != nil {
		return nil, ret.Error
	}

	return ret.Value.(*models.Shop), nil
}

// UpdateShopByWhere 更新shop
func (a *ShopRepository) UpdateShopByWhere(where interface{}, update interface{}) bool {
	if err := a.Base.Source.DB().Where(where).Update(update).Error; err != nil {
		a.Log.Errorf("更新失败", err)
		return false
	}
	return true
}
