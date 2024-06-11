package repository_impl

import (
	"hmdp-go/common/logger"
	"hmdp-go/models"
)

// SeckillVoucherRepository 注入IDb
type SeckillVoucherRepository struct {
	Log  logger.ILogger `inject:""`
	Base BaseRepository `inject:"inline"`
}

// GetTables 分页返回SeckillVouchers
func (a *SeckillVoucherRepository) GetTables(PageNum, PageSize int64, where interface{}) []*models.SeckillVoucher {
	var seckillVouchers []*models.SeckillVoucher
	var total uint64
	err := a.Base.GetPages(&models.SeckillVoucher{}, &seckillVouchers, PageNum, PageSize, &total, "")
	if err != nil {
		a.Log.Errorf("GetTables函数出错：", err)
	}
	return seckillVouchers
}

// GetSeckillVoucher 根据id获取SeckillVoucher
func (a *SeckillVoucherRepository) GetSeckillVoucher(where interface{}) *models.SeckillVoucher {
	var seckillVoucher models.SeckillVoucher
	if err := a.Base.First(where, &seckillVoucher); err != nil {
		a.Log.Errorf("未找到相关SeckillVoucher", err)
	}
	return &seckillVoucher
}

// AddSeckillVoucher 新增SeckillVoucher
func (a *SeckillVoucherRepository) AddSeckillVoucher(seckillVoucher *models.SeckillVoucher) bool {
	if err := a.Base.Save(seckillVoucher); err != nil {
		a.Log.Errorf("添加SeckillVoucher失败", err)
	}
	return true
}

// GetSeckillVouchers 获取SeckillVoucher
func (a *SeckillVoucherRepository) GetSeckillVouchers(PageNum int64, PageSize int64, total *uint64, where interface{}) []*models.SeckillVoucher {
	var seckillVouchers []*models.SeckillVoucher
	err := a.Base.GetPages(&models.SeckillVoucher{}, &seckillVouchers, PageNum, PageSize, total, where, "ID desc")
	if err != nil {
		a.Log.Errorf("获取SeckillVoucher信息失败", err)
	}
	return seckillVouchers
}
