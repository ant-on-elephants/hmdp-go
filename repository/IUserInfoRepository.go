package repository

import "hmdp-go/models"

// IUserInfoRepository UserInfo接口定义
type IUserInfoRepository interface {
	//GetTables 分页返回UserInfos
	GetTables(PageNum, PageSize int64, where interface{}) []*models.UserInfo
	//GetUserInfo 根据id获取UserInfo
	GetUserInfo(where interface{}) *models.UserInfo
	//AddUserInfo 新增UserInfo
	AddUserInfo(userInfo *models.UserInfo) bool
	//GetUserInfos 获取UserInfo
	GetUserInfos(PageNum int64, PageSize int64, total *uint64, where interface{}) []*models.UserInfo
}
