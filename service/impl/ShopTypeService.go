package impl

import (
	"hmdp-go/common/logger"
	"hmdp-go/models"
	"hmdp-go/repository"
)

// ShopTypeService 注入IDb
type ShopTypeService struct {
	Log          logger.ILogger                 `inject:""`
	ShopTypeRepo repository.IShopTypeRepository `inject:""`
}

func (s *ShopTypeService) GetShopTypeList(order interface{}) ([]*models.ShopType, error) {
	list, err := s.ShopTypeRepo.GetShopTypeList("", order)
	if err != nil {
		return nil, err
	}
	return list, nil
}
