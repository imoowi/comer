package impl

import (
	"github.com/gin-gonic/gin"
	"github.com/imoowi/comer/interfaces"
	"github.com/imoowi/comer/utils/response"
)

type Service struct {
	Repo *interfaces.IRepo
	MT   *interfaces.IModel
}

func NewService(r interfaces.IRepo) *Service {
	return &Service{
		Repo: &r,
	}
}
func (s *Service) PageList(c *gin.Context, filter *interfaces.IFilter, mt *interfaces.IModel) (res *response.PageList, err error) {
	repo := *s.Repo
	return repo.PageList(c, filter, mt)
}

func (s *Service) One(c *gin.Context, filter *interfaces.IFilter, id uint, mt *interfaces.IModel) (res *interfaces.IModel, err error) {
	repo := *s.Repo
	return repo.One(c, filter, id, mt)
}
func (s *Service) OneByName(c *gin.Context, filter *interfaces.IFilter, name string, mt *interfaces.IModel) (res *interfaces.IModel, err error) {
	repo := *s.Repo
	return repo.OneByName(c, filter, name, mt)
}
func (s *Service) Add(c *gin.Context, filter *interfaces.IFilter, model *interfaces.IModel, mt *interfaces.IModel) (newId uint, err error) {
	repo := *s.Repo
	return repo.Add(c, filter, model, mt)
}
func (s *Service) Update(c *gin.Context, filter *interfaces.IFilter, model *interfaces.IModel, id uint, mt *interfaces.IModel) (updated bool, err error) {
	repo := *s.Repo
	return repo.Update(c, filter, model, id, mt)
}
func (s *Service) Delete(c *gin.Context, filter *interfaces.IFilter, id uint, mt *interfaces.IModel) (deleted bool, err error) {
	repo := *s.Repo
	return repo.Delete(c, filter, id, mt)
}
