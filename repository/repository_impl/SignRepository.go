package repository_impl

import (
	"hmdp-go/common/logger"
	"hmdp-go/models"
)

// SignRepository 注入IDb
type SignRepository struct {
	Log  logger.ILogger `inject:""`
	Base BaseRepository `inject:"inline"`
}

// GetTables 分页返回Signs
func (a *SignRepository) GetTables(PageNum, PageSize int64, where interface{}) []*models.Sign {
	var signs []*models.Sign
	var total uint64
	err := a.Base.GetPages(&models.Sign{}, &signs, PageNum, PageSize, &total, "")
	if err != nil {
		a.Log.Errorf("GetTables函数出错：", err)
	}
	return signs
}

// GetSign 根据id获取Sign
func (a *SignRepository) GetSign(where interface{}) *models.Sign {
	var sign models.Sign
	if err := a.Base.First(where, &sign); err != nil {
		a.Log.Errorf("未找到相关Sign", err)
	}
	return &sign
}

// AddSign 新增Sign
func (a *SignRepository) AddSign(sign *models.Sign) bool {
	if err := a.Base.Save(sign); err != nil {
		a.Log.Errorf("添加Sign失败", err)
	}
	return true
}

// GetSigns 获取Sign
func (a *SignRepository) GetSigns(PageNum int64, PageSize int64, total *uint64, where interface{}) []*models.Sign {
	var signs []*models.Sign
	err := a.Base.GetPages(&models.Sign{}, &signs, PageNum, PageSize, total, where, "ID desc")
	if err != nil {
		a.Log.Errorf("获取Sign信息失败", err)
	}
	return signs
}
