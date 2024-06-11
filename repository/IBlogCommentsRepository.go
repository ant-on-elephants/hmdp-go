package repository

import "hmdp-go/models"

// IBlogCommentsRepository BlogComments接口定义
type IBlogCommentsRepository interface {
	//GetTables 分页返回BlogComments
	GetTables(PageNum, PageSize int64, where interface{}) []*models.BlogComments
	//GetBlogComment 根据id获取BlogComment
	GetBlogComment(where interface{}) *models.BlogComments
	//AddBlogComments 新增BlogComments
	AddBlogComments(blogComments *models.BlogComments) bool
	//GetBlogComments 获取BlogComments
	GetBlogComments(PageNum int64, PageSize int64, total *uint64, where interface{}) []*models.BlogComments
}
