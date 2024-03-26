package impl

import (
	"github.com/gin-gonic/gin"
	"github.com/imoowi/comer/utils/request"
	"gorm.io/gorm"
)

// 过滤器接口实现
type Filter struct {
	request.PageList
}

func (f *Filter) GetPage() int64 {
	return f.Page
}

func (f *Filter) SetPage(page int64) {
	f.Page = page
}

func (f *Filter) GetPageSize() int64 {
	return f.PageSize
}
func (f *Filter) SetPageSize(pageSize int64) {
	f.PageSize = pageSize
}

func (f *Filter) GetSearchKey() string {
	return f.SearchKey
}
func (f *Filter) SetSearchKey(searchKey string) {
	f.SearchKey = searchKey
}

// 分页查询过滤器构建方法
func (f *Filter) BuildPageListFilter(c *gin.Context, db *gorm.DB) *gorm.DB {
	return db
}
