package repository_impl

import (
	"hmdp-go/common/logger"
	"hmdp-go/models"
)

// ShopTypeRepository 注入IDb
type ShopTypeRepository struct {
	Log  logger.ILogger `inject:""`
	Base BaseRepository `inject:"inline"`
}

// GetTables 分页返回ShopTypes
func (a *ShopTypeRepository) GetTables(PageNum, PageSize int64, where interface{}) []*models.ShopType {
	var shopTypes []*models.ShopType
	var total uint64
	err := a.Base.GetPages(&models.ShopType{}, &shopTypes, PageNum, PageSize, &total, "")
	if err != nil {
		a.Log.Errorf("GetTables函数出错：", err)
	}
	return shopTypes
}

// GetShopType 根据id获取ShopType
func (a *ShopTypeRepository) GetShopType(where interface{}) *models.ShopType {
	var shopType models.ShopType
	if err := a.Base.First(where, &shopType); err != nil {
		a.Log.Errorf("未找到相关ShopType", err)
	}
	return &shopType
}

// AddShopType 新增ShopType
func (a *ShopTypeRepository) AddShopType(shopType *models.ShopType) bool {
	if err := a.Base.Save(shopType); err != nil {
		a.Log.Errorf("添加ShopType失败", err)
	}
	return true
}

// GetShopTypes 获取ShopType
func (a *ShopTypeRepository) GetShopTypes(PageNum int64, PageSize int64, total *uint64, where interface{}) []*models.ShopType {
	var shopTypes []*models.ShopType
	err := a.Base.GetPages(&models.ShopType{}, &shopTypes, PageNum, PageSize, total, where, "ID desc")
	if err != nil {
		a.Log.Errorf("获取ShopType信息失败", err)
	}
	return shopTypes
}

// GetShopTypeList 获取ShopType
func (a *ShopTypeRepository) GetShopTypeList(where interface{}, order interface{}) ([]*models.ShopType, error) {
	var shopTypes []*models.ShopType
	err := a.Base.Source.DB().Where(where).Order(order).Find(&shopTypes).Error
	if err != nil {
		a.Log.Errorf("获取ShopType信息失败", err)
		return nil, err
	}
	return shopTypes, nil
}
