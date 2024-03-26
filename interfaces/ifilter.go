package interfaces

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 过滤器接口
type IFilter interface {
	GetPage() int64
	SetPage(page int64)
	GetPageSize() int64
	SetPageSize(pageSize int64)
	GetSearchKey() string
	SetSearchKey(searchKey string)
	BuildPageListFilter(c *gin.Context, db *gorm.DB) *gorm.DB
}
