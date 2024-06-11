package models

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	"time"
)

type VoucherOrder struct {
	Id             int64        `db:"id"`              // 主键
	VoucherOrderId int64        `db:"voucherOrder_id"` // 下单的用户id
	VoucherId      int64        `db:"voucher_id"`      // 购买的代金券id
	PayType        int64        `db:"pay_type"`        // 支付方式 1：余额支付；2：支付宝；3：微信
	Status         int64        `db:"status"`          // 订单状态，1：未支付；2：已支付；3：已核销；4：已取消；5：退款中；6：已退款
	CreateTime     time.Time    `db:"create_time"`     // 下单时间
	PayTime        sql.NullTime `db:"pay_time"`        // 支付时间
	UseTime        sql.NullTime `db:"use_time"`        // 核销时间
	RefundTime     sql.NullTime `db:"refund_time"`     // 退款时间
	UpdateTime     time.Time    `db:"update_time"`     // 更新时间
}

// BeforeCreate 在创建VoucherOrder之前，先把创建时间赋值
func (voucherOrder *VoucherOrder) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("create_time", time.Now().Unix())
	return nil
}

// BeforeUpdate 在更新VoucherOrder之前，先把更新时间赋值
func (voucherOrder *VoucherOrder) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("update_time", time.Now().Unix())
	return nil
}
