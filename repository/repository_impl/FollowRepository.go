package repository_impl

import (
	"hmdp-go/common/logger"
	"hmdp-go/models"
)

// FollowRepository 注入IDb
type FollowRepository struct {
	Log  logger.ILogger `inject:""`
	Base BaseRepository `inject:"inline"`
}

// GetTables 分页返回Follows
func (a *FollowRepository) GetTables(PageNum, PageSize int64, where interface{}) []*models.Follow {
	var follows []*models.Follow
	var total uint64
	err := a.Base.GetPages(&models.Follow{}, &follows, PageNum, PageSize, &total, "")
	if err != nil {
		a.Log.Errorf("GetTables函数出错：", err)
	}
	return follows
}

// GetFollow 根据id获取Follow
func (a *FollowRepository) GetFollow(where interface{}) *models.Follow {
	var follow models.Follow
	if err := a.Base.First(where, &follow); err != nil {
		a.Log.Errorf("未找到相关Follow", err)
	}
	return &follow
}

// AddFollow 新增Follow
func (a *FollowRepository) AddFollow(follow *models.Follow) bool {
	if err := a.Base.Save(follow); err != nil {
		a.Log.Errorf("添加Follow失败", err)
	}
	return true
}

// GetFollows 获取Follow
func (a *FollowRepository) GetFollows(PageNum int64, PageSize int64, total *uint64, where interface{}) []*models.Follow {
	var follows []*models.Follow
	err := a.Base.GetPages(&models.Follow{}, &follows, PageNum, PageSize, total, where, "ID desc")
	if err != nil {
		a.Log.Errorf("获取Follow信息失败", err)
	}
	return follows
}
