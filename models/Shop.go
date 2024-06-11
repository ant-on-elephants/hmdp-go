package models

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	"time"
)

type Shop struct {
	Id         int64          `db:"id"`          // 主键
	Name       string         `db:"name"`        // 商铺名称
	TypeId     int64          `db:"type_id"`     // 商铺类型的id
	Images     string         `db:"images"`      // 商铺图片，多个图片以\',\'隔开
	Area       sql.NullString `db:"area"`        // 商圈，例如陆家嘴
	Address    string         `db:"address"`     // 地址
	X          float64        `db:"x"`           // 经度
	Y          float64        `db:"y"`           // 维度
	AvgPrice   sql.NullInt64  `db:"avg_price"`   // 均价，取整数
	Sold       int64          `db:"sold"`        // 销量
	Comments   int64          `db:"comments"`    // 评论数量
	Score      int64          `db:"score"`       // 评分，1~5分，乘10保存，避免小数
	OpenHours  sql.NullString `db:"open_hours"`  // 营业时间，例如 10:00-22:00
	CreateTime time.Time      `db:"create_time"` // 创建时间
	UpdateTime time.Time      `db:"update_time"` // 更新时间
}

// BeforeCreate 在创建Shop之前，先把创建时间赋值
func (shop *Shop) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("create_time", time.Now().Unix())
	return nil
}

// BeforeUpdate 在更新Shop之前，先把更新时间赋值
func (shop *Shop) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("update_time", time.Now().Unix())
	return nil
}
