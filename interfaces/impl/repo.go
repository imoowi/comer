package impl

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/imoowi/comer/components"
	"github.com/imoowi/comer/interfaces"
	"github.com/imoowi/comer/utils/response"
)

type Repo[T interfaces.IModel] struct {
	DB *components.MysqlODM
}

func NewRepo[T interfaces.IModel](db *components.MysqlODM) *Repo[T] {
	return &Repo[T]{
		DB: db,
	}
}
func (r *Repo[T]) PageList(c *gin.Context, f *interfaces.IFilter) (res *response.PageListT[T], err error) {
	db := r.DB.Client
	db = (*f).BuildPageListFilter(c, db)
	offset := ((*f).GetPage() - 1) * (*f).GetPageSize()
	db = db.Model(new(T)).Offset(int(offset)).Limit(int((*f).GetPageSize()))
	objs := make([]T, 0)
	err = db.Find(&objs).Error
	var count int64
	db.Offset(-1).Limit(-1).Count(&count)

	res = &response.PageListT[T]{
		List:  objs,
		Pages: response.MakePages(count, (*f).GetPage(), (*f).GetPageSize()),
	}

	return
}

func (r *Repo[T]) One(c *gin.Context, id uint) (res T, err error) {
	db := r.DB.Client
	err = db.Model(new(T)).Where(`id=?`, id).First(&res).Error
	return
}
func (r *Repo[T]) OneByName(c *gin.Context, name string) (res T, err error) {
	db := r.DB.Client
	err = db.Model(new(T)).Where(`name=?`, name).First(&res).Error
	return
}
func (r *Repo[T]) Add(c *gin.Context, model T) (newId uint, err error) {
	db := r.DB.Client
	err = db.Create(model).Error
	newId = model.GetID()
	return
}
func (r *Repo[T]) Update(c *gin.Context, updateFields map[string]any, id uint) (updated bool, err error) {
	if id <= 0 {
		updated = false
		err = errors.New(`pls input id`)
		return
	}
	_, err = r.One(c, id)
	if err != nil {
		return
	}
	db := r.DB.Client
	err = db.Model(new(T)).Omit(`created_at`).Where(`id=?`, id).Updates(updateFields).Error
	if err == nil {
		updated = true
	}
	return
}
func (r *Repo[T]) Delete(c *gin.Context, id uint) (deleted bool, err error) {
	if id <= 0 {
		deleted = false
		err = errors.New(`pls input id`)
		return
	}
	db := r.DB.Client
	model, err := r.One(c, id)
	if err != nil {
		return
	}
	err = db.Model(new(T)).Where(`id=?`, id).Delete(&model).Error
	if err == nil {
		deleted = true
	}
	return
}
