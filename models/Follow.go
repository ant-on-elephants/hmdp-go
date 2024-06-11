package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Follow struct {
	Id           int64     `db:"id"`             // 主键
	UserId       int64     `db:"user_id"`        // 用户id
	FollowUserId int64     `db:"follow_user_id"` // 关联的用户id
	CreateTime   time.Time `db:"create_time"`    // 创建时间
}

// BeforeCreate 在创建Follow之前，先把创建时间赋值
func (follow *Follow) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("create_time", time.Now().Unix())
	return nil
}
