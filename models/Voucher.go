package models

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	"time"
)

type Voucher struct {
	Id          int64          `db:"id"`           // 主键
	ShopId      sql.NullInt64  `db:"shop_id"`      // 商铺id
	Title       string         `db:"title"`        // 代金券标题
	SubTitle    sql.NullString `db:"sub_title"`    // 副标题
	Rules       sql.NullString `db:"rules"`        // 使用规则
	PayValue    int64          `db:"pay_value"`    // 支付金额，单位是分。例如200代表2元
	ActualValue int64          `db:"actual_value"` // 抵扣金额，单位是分。例如200代表2元
	Type        int64          `db:"type"`         // 0,普通券；1,秒杀券
	Status      int64          `db:"status"`       // 1,上架; 2,下架; 3,过期
	CreateTime  time.Time      `db:"create_time"`  // 创建时间
	UpdateTime  time.Time      `db:"update_time"`  // 更新时间
}

// BeforeCreate 在创建Voucher之前，先把创建时间赋值
func (voucher *Voucher) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("create_time", time.Now().Unix())
	return nil
}

// BeforeUpdate 在更新Voucher之前，先把更新时间赋值
func (voucher *Voucher) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("update_time", time.Now().Unix())
	return nil
}
