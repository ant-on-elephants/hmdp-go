package models

import (
	"database/sql"
	"time"
)

type Sign struct {
	Id       int64         `db:"id"`        // 主键
	UserId   int64         `db:"user_id"`   // 用户id
	Year     int64         `db:"year"`      // 签到的年
	Month    int64         `db:"month"`     // 签到的月
	Date     time.Time     `db:"date"`      // 签到的日期
	IsBackup sql.NullInt64 `db:"is_backup"` // 是否补签
}
