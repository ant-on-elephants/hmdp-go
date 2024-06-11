package models

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	"time"
)

type UserInfo struct {
	UserId     int64          `db:"user_id"`     // 主键，用户id
	City       string         `db:"city"`        // 城市名称
	Introduce  sql.NullString `db:"introduce"`   // 个人介绍，不要超过128个字符
	Fans       int64          `db:"fans"`        // 粉丝数量
	Followee   int64          `db:"followee"`    // 关注的人的数量
	Gender     int64          `db:"gender"`      // 性别，0：男，1：女
	Birthday   sql.NullTime   `db:"birthday"`    // 生日
	Credits    int64          `db:"credits"`     // 积分
	Level      int64          `db:"level"`       // 会员级别，0~9级,0代表未开通会员
	CreateTime time.Time      `db:"create_time"` // 创建时间
	UpdateTime time.Time      `db:"update_time"` // 更新时间
}

// BeforeCreate 在创建UserInfo之前，先把创建时间赋值
func (userInfo *UserInfo) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("create_time", time.Now().Unix())
	return nil
}

// BeforeUpdate 在更新UserInfo之前，先把更新时间赋值
func (userInfo *UserInfo) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("update_time", time.Now().Unix())
	return nil
}
