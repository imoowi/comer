/*
Copyright © 2023 yuanjun<simpleyuan@gmail.com>
*/
package repos

import (
	"errors"

	"github.com/gin-gonic/gin"
	"{{.ModuleName}}/apps/{{.appName}}/models"
	"{{.ModuleName}}/components"
	"{{.ModuleName}}/global"
	"{{.ModuleName}}/utils/request"
	"{{.ModuleName}}/utils/response"
)

type {{.ModelName}}Repo struct {
	Db *components.MysqlODM
}

func new{{.ModelName}}Repo() *{{.ModelName}}Repo {
	return &{{.ModelName}}Repo{
		Db: global.MysqlDb,
	}
}
func init() {
	global.DigContainer.Provide(new{{.ModelName}}Repo)
}
func (r *{{.ModelName}}Repo) PageList(c *gin.Context, req *request.PageList) (res *response.PageList, err error) {
	db := r.Db.Client
	var users []*models.{{.ModelName}}

	if req.SearchKey != `` {
		db = db.Where(`name LIKE ?`, `%`+req.SearchKey+`%`)
	}
	offset := (req.Page - 1) * req.PageSize
	db = db.Offset(int(offset)).Limit(int(req.PageSize))
	// db=db.Order(`name desc`)
	err = db.Find(&users).Error

	var count int64
	db.Offset(-1).Limit(-1).Count(&count)

	res = &response.PageList{
		List:  users,
		Pages: response.MakePages(count, req.Page, req.PageSize),
	}
	return
}

func (r *{{.ModelName}}Repo) One(c *gin.Context, id uint) (user *models.{{.ModelName}}, err error) {
	db := r.Db.Client
	err = db.Where(`id=?`, id).First(&user).Error
	return
}

func (r *{{.ModelName}}Repo) OneByName(c *gin.Context, name string) (user *models.{{.ModelName}}, err error) {
	db := r.Db.Client
	err = db.Where(`name=?`, name).First(&user).Error
	return
}
func (r *{{.ModelName}}Repo) Add(c *gin.Context, model *models.{{.ModelName}}) (newId uint, err error) {
	db := r.Db.Client
	db = db.Create(&model)
	err = db.Error

	newId =  model.ID
	return
}

func (r *{{.ModelName}}Repo) Update(c *gin.Context, model *models.{{.ModelName}}, id uint) (updated bool, err error) {
	if id == 0 {
		updated = false
		err = errors.New(`pls input id`)
		return
	}
	model.ID = uint(id)
	db := r.Db.Client
	err = db.Omit(`created_at`).Save(&model).Error
	if err == nil {
		updated = true
	}
	return
}

func (r *{{.ModelName}}Repo) Delete(c *gin.Context, id uint) (deleted bool, err error) {
	if id == 0 {
		deleted = false
		err = errors.New(`pls input id`)
		return
	}
	db := r.Db.Client
	model, err := r.One(c, id)
	if err != nil {
		return
	}
	if model.ID == 0 {
		err = errors.New(`obj not existe`)
		return
	}
	err = db.Delete(&model).Error
	if err == nil {
		deleted = true
	}
	return
}
