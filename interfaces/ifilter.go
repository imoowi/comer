package interfaces

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type IFilter interface {
	GetPage() int64
	SetPage(page int64)
	GetPageSize() int64
	SetPageSize(pageSize int64)
	GetSearchKey() string
	SetSearchKey(searchKey string)
	BuildPageListFilter(c *gin.Context, db *gorm.DB) *gorm.DB
	BuildOneFilter(c *gin.Context, db *gorm.DB) *gorm.DB
	BuildOneByNameFilter(c *gin.Context, db *gorm.DB) *gorm.DB
	BuildAddFilter(c *gin.Context, db *gorm.DB) *gorm.DB
	BuildUpdateFilter(c *gin.Context, db *gorm.DB) *gorm.DB
	BuildDelFilter(c *gin.Context, db *gorm.DB) *gorm.DB
}
