package impl

import (
	"github.com/gin-gonic/gin"
	"github.com/imoowi/comer/interfaces"
	"github.com/imoowi/comer/utils/response"
)

type Service[T interfaces.IModel] struct {
	Repo *interfaces.IRepo[T]
}

func NewService[T interfaces.IModel](r interfaces.IRepo[T]) *Service[T] {
	return &Service[T]{
		Repo: &r,
	}
}
func (s *Service[T]) PageList(c *gin.Context, filter *interfaces.IFilter) (res *response.PageListT[T], err error) {
	repo := *s.Repo
	return repo.PageList(c, filter)
}

func (s *Service[T]) One(c *gin.Context, id uint) (res T, err error) {
	repo := *s.Repo
	return repo.One(c, id)
}
func (s *Service[T]) OneByName(c *gin.Context, name string) (res T, err error) {
	repo := *s.Repo
	return repo.OneByName(c, name)
}
func (s *Service[T]) Add(c *gin.Context, model T) (newId uint, err error) {
	repo := *s.Repo
	return repo.Add(c, model)
}
func (s *Service[T]) Update(c *gin.Context, updateFields map[string]any, id uint) (updated bool, err error) {
	repo := *s.Repo
	return repo.Update(c, updateFields, id)
}
func (s *Service[T]) Delete(c *gin.Context, id uint) (deleted bool, err error) {
	repo := *s.Repo
	return repo.Delete(c, id)
}
