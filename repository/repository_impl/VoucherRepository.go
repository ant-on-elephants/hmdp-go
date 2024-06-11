package repository_impl

import (
	"hmdp-go/common/logger"
	"hmdp-go/models"
)

// VoucherRepository 注入IDb
type VoucherRepository struct {
	Log  logger.ILogger `inject:""`
	Base BaseRepository `inject:"inline"`
}

// GetTables 分页返回Vouchers
func (a *VoucherRepository) GetTables(PageNum, PageSize int64, where interface{}) []*models.Voucher {
	var vouchers []*models.Voucher
	var total uint64
	err := a.Base.GetPages(&models.Voucher{}, &vouchers, PageNum, PageSize, &total, "")
	if err != nil {
		a.Log.Errorf("GetTables函数出错：", err)
	}
	return vouchers
}

// GetVoucher 根据id获取Voucher
func (a *VoucherRepository) GetVoucher(where interface{}) *models.Voucher {
	var voucher models.Voucher
	if err := a.Base.First(where, &voucher); err != nil {
		a.Log.Errorf("未找到相关Voucher", err)
	}
	return &voucher
}

// AddVoucher 新增Voucher
func (a *VoucherRepository) AddVoucher(voucher *models.Voucher) bool {
	if err := a.Base.Save(voucher); err != nil {
		a.Log.Errorf("添加Voucher失败", err)
	}
	return true
}

// GetVouchers 获取Voucher
func (a *VoucherRepository) GetVouchers(PageNum int64, PageSize int64, total *uint64, where interface{}) []*models.Voucher {
	var vouchers []*models.Voucher
	err := a.Base.GetPages(&models.Voucher{}, &vouchers, PageNum, PageSize, total, where, "ID desc")
	if err != nil {
		a.Log.Errorf("获取Voucher信息失败", err)
	}
	return vouchers
}

// Save 保存
func (a *VoucherRepository) Save(voucher *models.Voucher) (*models.Voucher, error) {
	ret := a.Base.Source.DB().Save(voucher)
	if ret.Error != nil {
		return nil, ret.Error
	}

	return ret.Value.(*models.Voucher), nil
}
