package repository_impl

import (
	"hmdp-go/common/logger"
	"hmdp-go/models"
)

// BlogCommentsRepository 注入IDb
type BlogCommentsRepository struct {
	Log  logger.ILogger `inject:""`
	Base BaseRepository `inject:"inline"`
}

// GetTables 分页返回BlogComments
func (a *BlogCommentsRepository) GetTables(PageNum, PageSize int64, where interface{}) []*models.BlogComments {
	var blogComments []*models.BlogComments
	var total uint64
	err := a.Base.GetPages(&models.BlogComments{}, &blogComments, PageNum, PageSize, &total, "")
	if err != nil {
		a.Log.Errorf("GetTables函数出错：", err)
	}
	return blogComments
}

// GetBlogComment 根据id获取BlogComment
func (a *BlogCommentsRepository) GetBlogComment(where interface{}) *models.BlogComments {
	var blogComments models.BlogComments
	if err := a.Base.First(where, &blogComments); err != nil {
		a.Log.Errorf("未找到相关BlogComments", err)
	}
	return &blogComments
}

// AddBlogComments 新增BlogComment
func (a *BlogCommentsRepository) AddBlogComments(blogComments *models.BlogComments) bool {
	if err := a.Base.Save(blogComments); err != nil {
		a.Log.Errorf("添加BlogComments失败", err)
	}
	return true
}

// GetBlogComments 获取BlogComments
func (a *BlogCommentsRepository) GetBlogComments(PageNum int64, PageSize int64, total *uint64, where interface{}) []*models.BlogComments {
	var blogComments []*models.BlogComments
	err := a.Base.GetPages(&models.BlogComments{}, &blogComments, PageNum, PageSize, total, where, "ID desc")
	if err != nil {
		a.Log.Errorf("获取BlogComments信息失败", err)
	}
	return blogComments
}
