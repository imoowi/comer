package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/imoowi/comer/utils/response"
)

type IService interface {

	// 分页查询
	PageList(c *gin.Context, filter *IFilter, modelType *IModel) (res *response.PageList, err error)
	// 查询一个
	One(c *gin.Context, filter *IFilter, id uint, modelType *IModel) (res *IModel, err error)
	// 根据名称查询
	OneByName(c *gin.Context, filter *IFilter, name string, modelType *IModel) (res *IModel, err error)
	// 添加
	Add(c *gin.Context, filter *IFilter, model *IModel, modelType *IModel) (newId uint, err error)
	// 更新
	Update(c *gin.Context, filter *IFilter, model *IModel, id uint, modelType *IModel) (updated bool, err error)
	// 删除
	Delete(c *gin.Context, filter *IFilter, id uint, modelType *IModel) (deleted bool, err error)
}
