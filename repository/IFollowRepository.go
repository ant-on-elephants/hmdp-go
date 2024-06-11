package repository

import "hmdp-go/models"

// IFollowRepository Follow接口定义
type IFollowRepository interface {
	//GetTables 分页返回Follows
	GetTables(PageNum, PageSize int64, where interface{}) []*models.Follow
	//GetFollow 根据id获取Follow
	GetFollow(where interface{}) *models.Follow
	//AddFollow 新增Follow
	AddFollow(follow *models.Follow) bool
	//GetFollows 获取Follow
	GetFollows(PageNum int64, PageSize int64, total *uint64, where interface{}) []*models.Follow
}
