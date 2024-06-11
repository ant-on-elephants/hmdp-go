package repository_impl

import (
	"hmdp-go/common/logger"
	"hmdp-go/models"
)

// UserInfoRepository 注入IDb
type UserInfoRepository struct {
	Log  logger.ILogger `inject:""`
	Base BaseRepository `inject:"inline"`
}

// GetTables 分页返回UserInfos
func (a *UserInfoRepository) GetTables(PageNum, PageSize int64, where interface{}) []*models.UserInfo {
	var userInfos []*models.UserInfo
	var total uint64
	err := a.Base.GetPages(&models.UserInfo{}, &userInfos, PageNum, PageSize, &total, "")
	if err != nil {
		a.Log.Errorf("GetTables函数出错：", err)
	}
	return userInfos
}

// GetUserInfo 根据id获取UserInfo
func (a *UserInfoRepository) GetUserInfo(where interface{}) *models.UserInfo {
	var userInfo models.UserInfo
	if err := a.Base.First(where, &userInfo); err != nil {
		a.Log.Errorf("未找到相关UserInfo", err)
	}
	return &userInfo
}

// AddUserInfo 新增UserInfo
func (a *UserInfoRepository) AddUserInfo(userInfo *models.UserInfo) bool {
	if err := a.Base.Save(userInfo); err != nil {
		a.Log.Errorf("添加UserInfo失败", err)
	}
	return true
}

// GetUserInfos 获取UserInfo
func (a *UserInfoRepository) GetUserInfos(PageNum int64, PageSize int64, total *uint64, where interface{}) []*models.UserInfo {
	var userInfos []*models.UserInfo
	err := a.Base.GetPages(&models.UserInfo{}, &userInfos, PageNum, PageSize, total, where, "ID desc")
	if err != nil {
		a.Log.Errorf("获取UserInfo信息失败", err)
	}
	return userInfos
}
