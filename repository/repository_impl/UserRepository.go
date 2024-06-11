package repository_impl

import (
	"github.com/jinzhu/gorm"
	"hmdp-go/common/logger"
	"hmdp-go/models"
)

// UserRepository 注入IDb
type UserRepository struct {
	Log  logger.ILogger `inject:""`
	Base BaseRepository `inject:"inline"`
}

// CheckUser 身份验证
func (a *UserRepository) CheckUser(where interface{}) bool {
	var user models.User
	if err := a.Base.First(where, &user); err != nil {
		a.Log.Errorf("用户名或密码错误", err)
		return false
	}
	return true
}

// GetUsers 获取用户信息
func (a *UserRepository) GetUsers(PageNum int64, PageSize int64, total *uint64, where interface{}) []*models.User {
	var users []*models.User
	if err := a.Base.GetPages(&models.User{}, &users, PageNum, PageSize, total, where); err != nil {
		a.Log.Errorf("获取用户信息失败", err)
	}
	return users
}

// AddUser 新建用户
func (a *UserRepository) AddUser(user *models.User) bool {
	if err := a.Base.Create(user); err != nil {
		a.Log.Errorf("新建用户失败", err)
		return false
	}
	return true
}

// ExistUserByName 判断用户名是否已存在
func (a *UserRepository) ExistUserByName(where interface{}) bool {
	var user models.User
	sel := "id"
	err := a.Base.First(&where, &user, sel)
	//记录不存在错误(RecordNotFound)，返回false
	if gorm.IsRecordNotFoundError(err) {
		return false
	}
	//其他类型的错误，写下日志，返回false
	if err != nil {
		a.Log.Error(err)
		return false
	}
	return true
}

// GetUserByID 获取用户
func (a *UserRepository) GetUserByID(id int64) *models.User {
	var user models.User
	if err := a.Base.FirstByID(&user, id); err != nil {
		a.Log.Error(err)
	}
	return &user
}

// GetUserByWhere 获取用户
func (a *UserRepository) GetUserByWhere(where interface{}) *models.User {
	var user models.User
	if err := a.Base.Source.DB().Select("*").Where(where).First(&user); err.Error != nil {
		a.Log.Error(err)
	}
	return &user
}
