package impl

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/imoowi/comer/interfaces"
	"github.com/imoowi/comer/utils/response"
	"gorm.io/gorm"
)

// 数据资源接口实现
type Repo[T interfaces.IModel] struct {
	db *gorm.DB
}

// 新建一个数据资源
func NewRepo[T interfaces.IModel](db *gorm.DB) *Repo[T] {
	return &Repo[T]{
		db: db,
	}
}

// 分页查询数据
func (r *Repo[T]) PageList(c *gin.Context, f *interfaces.IFilter) (res *response.PageListT[T], err error) {
	db := r.db
	db = (*f).BuildPageListFilter(c, db)
	offset := ((*f).GetPage() - 1) * (*f).GetPageSize()
	db = db.Model(new(T)).Offset(int(offset)).Limit(int((*f).GetPageSize()))
	objs := make([]T, 0)
	err = db.Find(&objs).Error
	var count int64
	db.Offset(-1).Limit(-1).Select("count(id)").Count(&count)

	res = &response.PageListT[T]{
		List:  objs,
		Pages: response.MakePages(count, (*f).GetPage(), (*f).GetPageSize()),
	}

	return
}

// 分页查询数据
func (r *Repo[T]) PageListWithSelectOption(c *gin.Context, f *interfaces.IFilter, selectOpt []string) (res *response.PageListT[T], err error) {
	db := r.db
	db = (*f).BuildPageListFilter(c, db)
	offset := ((*f).GetPage() - 1) * (*f).GetPageSize()
	db = db.Model(new(T)).Offset(int(offset)).Limit(int((*f).GetPageSize()))
	if len(selectOpt) > 0 {
		db = db.Select(selectOpt)
	}
	objs := make([]T, 0)
	err = db.Find(&objs).Error
	var count int64
	db.Offset(-1).Limit(-1).Select("count(id)").Count(&count)

	res = &response.PageListT[T]{
		List:  objs,
		Pages: response.MakePages(count, (*f).GetPage(), (*f).GetPageSize()),
	}

	return
}

// 根据id查询一条记录
func (r *Repo[T]) One(c *gin.Context, id uint) (res T, err error) {
	db := r.db
	err = db.Model(new(T)).Where(`id=?`, id).First(&res).Error
	return
}

// 根据id查询一条记录
func (r *Repo[T]) OneWithSelectOption(c *gin.Context, id uint, selectOpt []string) (res T, err error) {
	db := r.db
	db = db.Model(new(T)).Where(`id=?`, id)
	if len(selectOpt) > 0 {
		db = db.Select(selectOpt)
	}
	err = db.First(&res).Error
	return
}

// 根据名字查询一条记录
func (r *Repo[T]) OneByName(c *gin.Context, name string) (res T, err error) {
	db := r.db
	err = db.Model(new(T)).Where(`name=?`, name).First(&res).Error
	return
}

// 根据名字查询一条记录
func (r *Repo[T]) OneByNameWithSelectOption(c *gin.Context, name string, selectOpt []string) (res T, err error) {
	db := r.db
	db = db.Model(new(T)).Where(`name=?`, name)
	if len(selectOpt) > 0 {
		db = db.Select(selectOpt)
	}
	err = db.First(&res).Error
	return
}

// 新建资源
func (r *Repo[T]) Add(c *gin.Context, model T) (newId uint, err error) {
	db := r.db
	err = db.Create(model).Error
	newId = model.GetID()
	return
}

// 通过id更新资源，只更新updateFields里有的字段
func (r *Repo[T]) Update(c *gin.Context, updateFields map[string]any, id uint) (updated bool, err error) {
	if id <= 0 {
		updated = false
		err = errors.New(`pls input id`)
		return
	}
	_model, err := r.One(c, id)
	if err != nil {
		return
	}
	db := r.db
	err = db.Model(_model).Omit(`created_at`).Where(`id=?`, id).Updates(updateFields).Error
	if err == nil {
		updated = true
	}
	return
}

// 根据id删除资源
func (r *Repo[T]) Delete(c *gin.Context, id uint) (deleted bool, err error) {
	if id <= 0 {
		deleted = false
		err = errors.New(`pls input id`)
		return
	}
	db := r.db
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
