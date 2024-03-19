package models

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"github.com/imoowi/comer/interfaces/impl"
)

// 查询参数
type UserFilter struct {
	impl.Filter
}

func (f *UserFilter) BuildPageListFilter(c *gin.Context, db *gorm.DB) *gorm.DB {
	if f.GetSearchKey() != `` {
		db = db.Where(`name LIKE ?`, `%`+f.GetSearchKey()+`%`)
	}
	return f.Filter.BuildPageListFilter(c, db)
}
