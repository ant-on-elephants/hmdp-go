package repository

import "hmdp-go/models"

// IVoucherRepository Voucher接口定义
type IVoucherRepository interface {
	//GetTables 分页返回Vouchers
	GetTables(PageNum, PageSize int64, where interface{}) []*models.Voucher
	//GetVoucher 根据id获取Voucher
	GetVoucher(where interface{}) *models.Voucher
	//AddVoucher 新增Voucher
	AddVoucher(voucher *models.Voucher) bool
	//GetVouchers 获取Voucher
	GetVouchers(PageNum int64, PageSize int64, total *uint64, where interface{}) []*models.Voucher

	Save(voucher *models.Voucher) (*models.Voucher, error)
}
