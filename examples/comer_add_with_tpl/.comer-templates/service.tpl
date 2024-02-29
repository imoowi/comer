// generated by "comer add-with-tpl"
// © 2024 IMOOWI. All Rights Reserved.
//
//	@Author	yuanjun<imoowi@qq.com>
package services

import (
	"{{.ModuleName}}/internal/global"
	"{{.ModuleName}}/internal/models"
	"{{.ModuleName}}/internal/util/response"
	"{{.ModuleName}}/pkg/frame"
)

var {{.ServiceName}} *{{.ServiceName}}Service

func Register{{.ServiceName}}Service(s *{{.ServiceName}}Service) {
	{{.ServiceName}} = s
}

type {{.ServiceName}}Service struct {
	{{.ModelName}}Repo *models.{{.ModelName}}Repo
}

func new{{.ServiceName}}Service(r *models.{{.ModelName}}Repo) *{{.ServiceName}}Service {
	return &{{.ServiceName}}Service{
		{{.ModelName}}Repo: r,
	}
}

func init() {
	global.Container.Provide(new{{.ServiceName}}Service)
	global.RegisterContainerProviders(Register{{.ServiceName}}Service)
}

func (s *{{.ServiceName}}Service) All(c *frame.Context, query *models.{{.ModelName}}Query) (res []*models.{{.ModelName}}, err error) {
	res, err = s.{{.ModelName}}Repo.All(c, query)
	return
}
func (s *{{.ServiceName}}Service) PageList(c *frame.Context, req *models.{{.ModelName}}Query) (res *response.PageList, err error) {
	res, err = s.{{.ModelName}}Repo.PageList(c, req)
	return
}

func (s *{{.ServiceName}}Service) One(c *frame.Context, id uint) (model *models.{{.ModelName}}, err error) {
	model, err = s.{{.ModelName}}Repo.One(c, id)
	return
}

func (s *{{.ServiceName}}Service) OneByName(c *frame.Context, name string) (model *models.{{.ModelName}}, err error) {
	model, err = s.{{.ModelName}}Repo.OneByName(c, name)
	return
}

func (s *{{.ServiceName}}Service) Add(c *frame.Context, _model *models.{{.ModelName}}) (newId uint, err error) {
	// model, err := s.{{.ModelName}}Repo.OneByName(c, _model.Name)
	// if model != nil && model.ID > 0 {
	// 	newId = 0
	// 	err = errors.New(`name existed`)
	// 	return
	// }
	_model.ProjectId = c.GetUint(`project_id`)
	newId, err = s.{{.ModelName}}Repo.Add(c, _model)
	return
}

func (s *{{.ServiceName}}Service) Update(c *frame.Context, _model *models.{{.ModelName}}, id uint) (updated bool, err error) {
	_model.ProjectId = c.GetUint(`project_id`)
	updated, err = s.{{.ModelName}}Repo.Update(c, _model, id)
	return
}
func (s *{{.ServiceName}}Service) Delete(c *frame.Context, id uint) (deleted bool, err error) {
	deleted, err = s.{{.ModelName}}Repo.Delete(c, id)
	return
}
