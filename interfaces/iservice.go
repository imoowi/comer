package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/imoowi/comer/utils/response"
)

// 服务层接口
type IService[T IModel] interface {

	// 分页查询
	PageList(c *gin.Context, filter *IFilter) (res *response.PageListT[T], err error)
	// 查询一个
	One(c *gin.Context, id uint) (res T, err error)
	// 根据名称查询
	OneByName(c *gin.Context, name string) (res T, err error)
	// 添加
	Add(c *gin.Context, model T) (newId uint, err error)
	// 更新,传什么就更新什么
	Update(c *gin.Context, updateFields map[string]any, id uint) (updated bool, err error)
	// 删除
	Delete(c *gin.Context, id uint) (deleted bool, err error)
}
