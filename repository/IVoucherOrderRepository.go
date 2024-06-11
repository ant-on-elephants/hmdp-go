package repository

import "hmdp-go/models"

// IVoucherOrderRepository VoucherOrder接口定义
type IVoucherOrderRepository interface {
	//GetTables 分页返回VoucherOrders
	GetTables(PageNum, PageSize int64, where interface{}) []*models.VoucherOrder
	//GetVoucherOrder 根据id获取VoucherOrder
	GetVoucherOrder(where interface{}) *models.VoucherOrder
	//AddVoucherOrder 新增VoucherOrder
	AddVoucherOrder(voucherOrder *models.VoucherOrder) bool
	//GetVoucherOrders 获取VoucherOrder
	GetVoucherOrders(PageNum int64, PageSize int64, total *uint64, where interface{}) []*models.VoucherOrder
}
