package repository_impl

import (
	"hmdp-go/common/logger"
	"hmdp-go/models"
)

// VoucherOrderRepository 注入IDb
type VoucherOrderRepository struct {
	Log  logger.ILogger `inject:""`
	Base BaseRepository `inject:"inline"`
}

// GetTables 分页返回VoucherOrders
func (a *VoucherOrderRepository) GetTables(PageNum, PageSize int64, where interface{}) []*models.VoucherOrder {
	var voucherOrders []*models.VoucherOrder
	var total uint64
	err := a.Base.GetPages(&models.VoucherOrder{}, &voucherOrders, PageNum, PageSize, &total, "")
	if err != nil {
		a.Log.Errorf("GetTables函数出错：", err)
	}
	return voucherOrders
}

// GetVoucherOrder 根据id获取VoucherOrder
func (a *VoucherOrderRepository) GetVoucherOrder(where interface{}) *models.VoucherOrder {
	var voucherOrder models.VoucherOrder
	if err := a.Base.First(where, &voucherOrder); err != nil {
		a.Log.Errorf("未找到相关VoucherOrder", err)
	}
	return &voucherOrder
}

// AddVoucherOrder 新增VoucherOrder
func (a *VoucherOrderRepository) AddVoucherOrder(voucherOrder *models.VoucherOrder) bool {
	if err := a.Base.Save(voucherOrder); err != nil {
		a.Log.Errorf("添加VoucherOrder失败", err)
	}
	return true
}

// GetVoucherOrders 获取VoucherOrder
func (a *VoucherOrderRepository) GetVoucherOrders(PageNum int64, PageSize int64, total *uint64, where interface{}) []*models.VoucherOrder {
	var voucherOrders []*models.VoucherOrder
	err := a.Base.GetPages(&models.VoucherOrder{}, &voucherOrders, PageNum, PageSize, total, where, "ID desc")
	if err != nil {
		a.Log.Errorf("获取VoucherOrder信息失败", err)
	}
	return voucherOrders
}
