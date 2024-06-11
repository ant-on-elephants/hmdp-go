package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type SeckillVoucher struct {
	VoucherId  int64     `db:"voucher_id"`  // 关联的优惠券的id
	Stock      int64     `db:"stock"`       // 库存
	CreateTime time.Time `db:"create_time"` // 创建时间
	BeginTime  time.Time `db:"begin_time"`  // 生效时间
	EndTime    time.Time `db:"end_time"`    // 失效时间
	UpdateTime time.Time `db:"update_time"` // 更新时间
}

// BeforeCreate 在创建SeckillVoucher之前，先把创建时间赋值
func (seckillVoucher *SeckillVoucher) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("create_time", time.Now().Unix())
	return nil
}

// BeforeUpdate 在更新SeckillVoucher之前，先把更新时间赋值
func (seckillVoucher *SeckillVoucher) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("update_time", time.Now().Unix())
	return nil
}
