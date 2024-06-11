package models

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	"time"
)

type BlogComments struct {
	Id         int64         `db:"id"`          // 主键
	UserId     int64         `db:"user_id"`     // 用户id
	BlogId     int64         `db:"blog_id"`     // 探店id
	ParentId   int64         `db:"parent_id"`   // 关联的1级评论id，如果是一级评论，则值为0
	AnswerId   int64         `db:"answer_id"`   // 回复的评论id
	Content    string        `db:"content"`     // 回复的内容
	Liked      sql.NullInt64 `db:"liked"`       // 点赞数
	Status     sql.NullInt64 `db:"status"`      // 状态，0：正常，1：被举报，2：禁止查看
	CreateTime time.Time     `db:"create_time"` // 创建时间
	UpdateTime time.Time     `db:"update_time"` // 更新时间
}

// BeforeCreate 在创建BlogComments之前，先把创建时间赋值
func (blogComments *BlogComments) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("create_time", time.Now().Unix())
	return nil
}

// BeforeUpdate 在更新Blog之前，先把更新时间赋值
func (blogComments *BlogComments) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("update_time", time.Now().Unix())
	return nil
}
