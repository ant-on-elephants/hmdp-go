package repository

import "hmdp-go/models"

// ISeckillVoucherRepository SeckillVoucher接口定义
type ISeckillVoucherRepository interface {
	//GetTables 分页返回SeckillVouchers
	GetTables(PageNum, PageSize int64, where interface{}) []*models.SeckillVoucher
	//GetSeckillVoucher 根据id获取SeckillVoucher
	GetSeckillVoucher(where interface{}) *models.SeckillVoucher
	//AddSeckillVoucher 新增SeckillVoucher
	AddSeckillVoucher(seckillVoucher *models.SeckillVoucher) bool
	//GetSeckillVouchers 获取SeckillVoucher
	GetSeckillVouchers(PageNum int64, PageSize int64, total *uint64, where interface{}) []*models.SeckillVoucher
}
