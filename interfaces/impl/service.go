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

func (s *Service[T]) One(c *gin.Context, filter *interfaces.IFilter, id uint) (res T, err error) {
	repo := *s.Repo
	return repo.One(c, filter, id)
}
func (s *Service[T]) OneByName(c *gin.Context, filter *interfaces.IFilter, name string) (res T, err error) {
	repo := *s.Repo
	return repo.OneByName(c, filter, name)
}
func (s *Service[T]) Add(c *gin.Context, filter *interfaces.IFilter, model T) (newId uint, err error) {
	repo := *s.Repo
	return repo.Add(c, filter, model)
}
func (s *Service[T]) Update(c *gin.Context, filter *interfaces.IFilter, model T, id uint) (updated bool, err error) {
	repo := *s.Repo
	return repo.Update(c, filter, model, id)
}
func (s *Service[T]) Delete(c *gin.Context, filter *interfaces.IFilter, id uint) (deleted bool, err error) {
	repo := *s.Repo
	return repo.Delete(c, filter, id)
}
