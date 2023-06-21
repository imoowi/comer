/*
Copyright © 2023 yuanjun<simpleyuan@gmail.com>
*/
package services

import (
	"errors"

	"github.com/gin-gonic/gin"
	"{{.ModuleName}}/apps/{{.appName}}/models"
	"{{.ModuleName}}/apps/{{.appName}}/repos"
	"{{.ModuleName}}/global"
	"{{.ModuleName}}/utils/request"
	"{{.ModuleName}}/utils/response"
)

var {{.ServiceName}} *{{.ServiceName}}Service

func Register{{.ServiceName}}Service(s *{{.ServiceName}}Service) {
	{{.ServiceName}} = s
}

type {{.ServiceName}}Service struct {
	{{.ModelName}}Repo *repos.{{.ModelName}}Repo
}

func new{{.ServiceName}}Service(r *repos.{{.ModelName}}Repo) *{{.ServiceName}}Service {
	return &{{.ServiceName}}Service{
		{{.ModelName}}Repo: r,
	}
}

func init() {
	global.DigContainer.Provide(new{{.ServiceName}}Service)
	global.RegisterContainerProviders(Register{{.ServiceName}}Service)
}

func (s *{{.ServiceName}}Service) PageList(c *gin.Context, req *request.PageList) (res *response.PageList, err error) {
	res, err = s.{{.ModelName}}Repo.PageList(c, req)
	return
}

func (s *{{.ServiceName}}Service) One(c *gin.Context, id uint) (model *models.{{.ModelName}}, err error) {
	model, err = s.{{.ModelName}}Repo.One(c, id)
	return
}

func (s *{{.ServiceName}}Service) OneByName(c *gin.Context, name string) (model *models.{{.ModelName}}, err error) {
	model, err = s.{{.ModelName}}Repo.OneByName(c, name)
	return
}

func (s *{{.ServiceName}}Service) Add(c *gin.Context, _model *models.{{.ModelName}}) (newId uint, err error) {
	model, err := s.{{.ModelName}}Repo.OneByName(c, _model.Name)
	if model != nil && model.ID > 0 {
		newId = 0
		err = errors.New(`name existed`)
		return
	}
	newId, err = s.{{.ModelName}}Repo.Add(c, _model)
	return
}

func (s *{{.ServiceName}}Service) Update(c *gin.Context, _model *models.{{.ModelName}}, id uint) (updated bool, err error) {
	updated, err = s.{{.ModelName}}Repo.Update(c, _model, id)
	return
}
func (s *{{.ServiceName}}Service) Delete(c *gin.Context, id uint) (deleted bool, err error) {
	deleted, err = s.{{.ModelName}}Repo.Delete(c, id)
	return
}