package repository_impl

import (
	"hmdp-go/common/logger"
	"hmdp-go/models"
)

// BlogRepository 注入IDb
type BlogRepository struct {
	Log  logger.ILogger `inject:""`
	Base BaseRepository `inject:"inline"`
}

// GetTables 分页返回Blogs
func (a *BlogRepository) GetTables(PageNum, PageSize int64, where interface{}) []*models.Blog {
	var blogs []*models.Blog
	var total uint64
	err := a.Base.GetPages(&models.Blog{}, &blogs, PageNum, PageSize, &total, "")
	if err != nil {
		a.Log.Errorf("GetTables函数出错：", err)
	}
	return blogs
}

// GetBlog 根据id获取Blog
func (a *BlogRepository) GetBlog(where interface{}) *models.Blog {
	var blog models.Blog
	if err := a.Base.First(where, &blog); err != nil {
		a.Log.Errorf("未找到相关Blog", err)
	}
	return &blog
}

// AddBlog 新增Blog
func (a *BlogRepository) AddBlog(blog *models.Blog) bool {
	if err := a.Base.Save(blog); err != nil {
		a.Log.Errorf("添加Blog失败", err)
	}
	return true
}

// GetBlogs 获取Blog
func (a *BlogRepository) GetBlogs(PageNum int64, PageSize int64, total *uint64, where interface{}) []*models.Blog {
	var blogs []*models.Blog
	err := a.Base.GetPages(&models.Blog{}, &blogs, PageNum, PageSize, total, where, "ID desc")
	if err != nil {
		a.Log.Errorf("获取Blog信息失败", err)
	}
	return blogs
}

// UpdateBlog 更新blog
func (a *BlogRepository) UpdateBlog(blog *models.Blog) bool {
	if err := a.Base.Source.DB().Update(blog).Error; err != nil {
		a.Log.Errorf("更新失败", err)
		return false
	}
	return true
}

// UpdateBlogByWhere 更新blog
func (a *BlogRepository) UpdateBlogByWhere(where interface{}, update interface{}) bool {
	if err := a.Base.Source.DB().Where(where).Update(update).Error; err != nil {
		a.Log.Errorf("更新失败", err)
		return false
	}
	return true
}
