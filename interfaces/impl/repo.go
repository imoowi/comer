package impl

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/imoowi/comer/components"
	"github.com/imoowi/comer/interfaces"
	"github.com/imoowi/comer/utils/response"
)

type Repo struct {
	DB *components.MysqlODM
}

func NewRepo(db *components.MysqlODM) *Repo {
	return &Repo{
		DB: db,
	}
}
func (r *Repo) PageList(c *gin.Context, q *interfaces.IFilter, modelType *interfaces.IModel) (res *response.PageList, err error) {
	db := r.DB.Client
	db = (*q).BuildPageListFilter(c, db)
	offset := ((*q).GetPage() - 1) * (*q).GetPageSize()
	db = db.Model(modelType).Offset(int(offset)).Limit(int((*q).GetPageSize()))
	objs := make([]map[string]any, 0)
	err = db.Find(&objs).Error
	//	if errors.Is(err, gorm.ErrRecordNotFound) {
	//		return
	//	}

	var count int64
	db.Offset(-1).Limit(-1).Count(&count)

	res = &response.PageList{
		List:  objs,
		Pages: response.MakePages(count, (*q).GetPage(), (*q).GetPageSize()),
	}

	return
}

func (r *Repo) One(c *gin.Context, q *interfaces.IFilter, id uint, modelType *interfaces.IModel) (res *interfaces.IModel, err error) {
	db := r.DB.Client
	db = (*q).BuildOneFilter(c, db)
	err = db.Model(modelType).Where(`id=?`, id).First(&modelType).Error
	// if errors.Is(err, gorm.ErrRecordNotFound) {
	// 	return
	// }
	res = modelType
	return
}
func (r *Repo) OneByName(c *gin.Context, q *interfaces.IFilter, name string, modelType *interfaces.IModel) (res *interfaces.IModel, err error) {
	db := r.DB.Client
	db = (*q).BuildOneByNameFilter(c, db)
	var obj *interfaces.IModel = modelType

	err = db.Model(modelType).First(obj).Error
	// if errors.Is(err, gorm.ErrRecordNotFound) {
	// 	return
	// }
	res = obj
	return
}
func (r *Repo) Add(c *gin.Context, q *interfaces.IFilter, model *interfaces.IModel, modelType *interfaces.IModel) (newId uint, err error) {
	db := r.DB.Client
	db = (*q).BuildAddFilter(c, db)
	var _model interfaces.IModel = *model
	db = db.Model(modelType).Create(_model)
	err = db.Error
	newId = (*model).GetID()
	return
}
func (r *Repo) Update(c *gin.Context, q *interfaces.IFilter, model *interfaces.IModel, id uint, modelType *interfaces.IModel) (updated bool, err error) {
	if id <= 0 {
		updated = false
		err = errors.New(`pls input id`)
		return
	}
	db := r.DB.Client
	db = (*q).BuildUpdateFilter(c, db)
	(*model).SetId(id)
	err = db.Model(modelType).Omit(`created_at`).Where(`id=?`, id).Save(model).Error
	if err == nil {
		updated = true
	}
	return
}
func (r *Repo) Delete(c *gin.Context, q *interfaces.IFilter, id uint, modelType *interfaces.IModel) (deleted bool, err error) {
	if id <= 0 {
		deleted = false
		err = errors.New(`pls input id`)
		return
	}
	db := r.DB.Client
	db = (*q).BuildDelFilter(c, db)
	model, err := r.One(c, q, id, modelType)
	if err != nil {
		return
	}
	err = db.Model(modelType).Where(`id=?`, id).Delete(&model).Error
	if err == nil {
		deleted = true
	}
	return
}
