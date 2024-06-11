package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	Id         int64     `db:"id"`          // 主键
	Phone      string    `db:"phone"`       // 手机号码
	Password   string    `db:"password"`    // 密码，加密存储
	NickName   string    `db:"nick_name"`   // 昵称，默认是用户id
	Icon       string    `db:"icon"`        // 人物头像
	CreateTime time.Time `db:"create_time"` // 创建时间
	UpdateTime time.Time `db:"update_time"` // 更新时间
}

// BeforeCreate 在创建User之前，先把创建时间赋值
func (user *User) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("create_time", time.Now().Unix())
	return nil
}

// BeforeUpdate 在更新User之前，先把更新时间赋值
func (user *User) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("update_time", time.Now().Unix())
	return nil
}
