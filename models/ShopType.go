package models

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	"time"
)

type ShopType struct {
	Id         int64          `db:"id"`          // 主键
	Name       sql.NullString `db:"name"`        // 类型名称
	Icon       sql.NullString `db:"icon"`        // 图标
	Sort       sql.NullInt64  `db:"sort"`        // 顺序
	CreateTime time.Time      `db:"create_time"` // 创建时间
	UpdateTime time.Time      `db:"update_time"` // 更新时间
}

// BeforeCreate 在创建ShopType之前，先把创建时间赋值
func (shopType *ShopType) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("create_time", time.Now().Unix())
	return nil
}

// BeforeUpdate 在更新ShopType之前，先把更新时间赋值
func (shopType *ShopType) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("update_time", time.Now().Unix())
	return nil
}
