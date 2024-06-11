package repository

import "hmdp-go/models"

// IBlogRepository Blog接口定义
type IBlogRepository interface {
	//GetTables 分页返回Blogs
	GetTables(PageNum, PageSize int64, where interface{}) []*models.Blog
	//GetBlog 根据id获取Blog
	GetBlog(where interface{}) *models.Blog
	//AddBlog 新增Blog
	AddBlog(blog *models.Blog) bool
	//GetBlogs 获取Blog
	GetBlogs(PageNum int64, PageSize int64, total *uint64, where interface{}) []*models.Blog
	//UpdateBlog 更新Blog
	UpdateBlog(blog *models.Blog) bool
	//UpdateBlogByWhere 更新Blog
	UpdateBlogByWhere(where interface{}, update interface{}) bool
}
