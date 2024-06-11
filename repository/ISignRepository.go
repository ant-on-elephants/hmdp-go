package repository

import "hmdp-go/models"

// ISignRepository Sign接口定义
type ISignRepository interface {
	//GetTables 分页返回Signs
	GetTables(PageNum, PageSize int64, where interface{}) []*models.Sign
	//GetSign 根据id获取Sign
	GetSign(where interface{}) *models.Sign
	//AddSign 新增Sign
	AddSign(sign *models.Sign) bool
	//GetSigns 获取Sign
	GetSigns(PageNum int64, PageSize int64, total *uint64, where interface{}) []*models.Sign
}
