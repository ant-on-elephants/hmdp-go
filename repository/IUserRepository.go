package repository

import "hmdp-go/models"

// IUserRepository User接口定义
type IUserRepository interface {
	//CheckUser 身份验证
	CheckUser(where interface{}) bool
	//GetUsers 获取用户信息
	GetUsers(PageNum int64, PageSize int64, total *uint64, where interface{}) []*models.User
	//AddUser 新建用户
	AddUser(user *models.User) bool
	//ExistUserByName 判断用户名是否已存在
	ExistUserByName(where interface{}) bool
	//GetUserByID 获取用户
	GetUserByID(id int64) *models.User
	//GetUserByWhere 获取用户
	GetUserByWhere(where interface{}) *models.User
}
