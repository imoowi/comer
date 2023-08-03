/*
generated by comer,https://github.com/imoowi/comer
Copyright © 2023 jun<simpleyuan@gmail.com>
*/
package repos

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin"
	"{{.ModuleName}}/apps/{{.appName}}/models"
	"{{.ModuleName}}/components"
	"{{.ModuleName}}/global"
	"{{.ModuleName}}/utils/response"
	"gorm.io/gorm"
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


func (r *{{.ModelName}}Repo) All(c *gin.Context, query *models.{{.ModelName}}Query) (res []*models.{{.ModelName}}, err error) {
	db := r.Db.Client
	if query.Key != `` {
		db = db.Where("`key`=?", query.Key)
	}

	err = db.Find(&res).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = nil
	}
	return
}

func (r *{{.ModelName}}Repo) PageList(c *gin.Context, query *models.{{.ModelName}}Query) (res *response.PageList, err error) {
	db := r.Db.Client
	var users []*models.{{.ModelName}}

	if query.SearchKey != `` {
		db = db.Where(`name LIKE ?`, `%`+query.SearchKey+`%`)
	}
	offset := (query.Page - 1) * query.PageSize
	db = db.Offset(int(offset)).Limit(int(query.PageSize))
	// db=db.Order(`name desc`)
	err = db.Find(&users).Error

	var count int64
	db.Offset(-1).Limit(-1).Count(&count)

	res = &response.PageList{
		List:  users,
		Pages: response.MakePages(count, query.Page, query.PageSize),
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


func (r *{{.ModelName}}Repo) PatchUpdate(c *gin.Context, patchData map[string]any, id uint) (updated bool, err error) {
	if id == 0 {
		updated = false
		err = errors.New(`pls input id`)
		return
	}
	model, err := r.One(c, id)
	if err != nil {
		return
	}
	if model == nil {
		err = errors.New(`no data existed`)
		return
	}

	patchDataBytes, err := json.Marshal(patchData)
	if err != nil {
		return
	}
	err = json.Unmarshal(patchDataBytes, &model)
	if err != nil {
		return
	}

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
