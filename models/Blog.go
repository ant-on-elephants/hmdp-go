package models

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	"time"
)

type Blog struct {
	Id         int64         `db:"id"`          // 主键
	ShopId     int64         `db:"shop_id"`     // 商户id
	UserId     int64         `db:"user_id"`     // 用户id
	Icon       string        `db:"-"`           // 用户图标
	Name       string        `db:"-"`           // 用户姓名
	IsLike     bool          `db:"-"`           // 是否点赞过了
	Title      string        `db:"title"`       // 标题
	Images     string        `db:"images"`      // 探店的照片，最多9张，多张以\",\"隔开
	Content    string        `db:"content"`     // 探店的文字描述
	Liked      int64         `db:"liked"`       // 点赞数量
	Comments   sql.NullInt64 `db:"comments"`    // 评论数量
	CreateTime time.Time     `db:"create_time"` // 创建时间
	UpdateTime time.Time     `db:"update_time"` // 更新时间
}

// BeforeCreate 在创建Blog之前，先把创建时间赋值
func (blog *Blog) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("create_time", time.Now().Unix())
	return nil
}

// BeforeUpdate 在更新Blog之前，先把更新时间赋值
func (blog *Blog) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("update_time", time.Now().Unix())
	return nil
}
