package impl

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/imoowi/comer/interfaces"
	"github.com/imoowi/comer/utils/response"
	"gorm.io/gorm"
)

// 服务层接口实现
type Service[T interfaces.IModel] struct {
	Repo *interfaces.IRepo[T]
}

// 新建服务
func NewService[T interfaces.IModel](r interfaces.IRepo[T]) *Service[T] {
	return &Service[T]{
		Repo: &r,
	}
}

// 分页查询
func (s *Service[T]) PageList(c *gin.Context, filter *interfaces.IFilter) (res *response.PageListT[T], err error) {
	repo := *s.Repo
	return repo.PageList(c, filter)
}

// 分页查询
func (s *Service[T]) PageListWithSelectOption(c *gin.Context, filter *interfaces.IFilter, selectOpt []string) (res *response.PageListT[T], err error) {
	repo := *s.Repo
	return repo.PageListWithSelectOption(c, filter, selectOpt)
}

// 查一条，根据id
func (s *Service[T]) One(c *gin.Context, id uint) (res T, err error) {
	repo := *s.Repo
	return repo.One(c, id)
}

// 查一条，根据id
func (s *Service[T]) OneWithSelectOption(c *gin.Context, id uint, selectOpt []string) (res T, err error) {
	repo := *s.Repo
	return repo.OneWithSelectOption(c, id, selectOpt)
}

// 查一条，根据名字
func (s *Service[T]) OneByName(c *gin.Context, name string) (res T, err error) {
	repo := *s.Repo
	return repo.OneByName(c, name)
}

// 查一条，根据名字
func (s *Service[T]) OneByNameWithSelectOption(c *gin.Context, name string, selectOpt []string) (res T, err error) {
	repo := *s.Repo
	return repo.OneByNameWithSelectOption(c, name, selectOpt)
}

// 新建资源
func (s *Service[T]) Add(c *gin.Context, model T) (newId uint, err error) {
	repo := *s.Repo
	return repo.Add(c, model)
}

// 更新资源
func (s *Service[T]) Update(c *gin.Context, updateFields map[string]any, id uint) (updated bool, err error) {
	delete(updateFields, `created_at`)
	delete(updateFields, `updated_at`)
	delete(updateFields, `deleted_at`)
	repo := *s.Repo
	return repo.Update(c, updateFields, id)
}

// 删除资源，根据id
func (s *Service[T]) Delete(c *gin.Context, id uint) (deleted bool, err error) {
	repo := *s.Repo
	return repo.Delete(c, id)
}

// extendedRepo 是 Repo[T] 的扩展方法集合，Service 通过类型断言访问（不改 IRepo 接口）。
type extendedRepo[T interfaces.IModel] interface {
	Transaction(fn func(tx *gorm.DB) error) error
	BatchAdd(c *gin.Context, models []T) error
	BatchDelete(c *gin.Context, ids []uint) error
	UnscopedOne(c *gin.Context, id uint) (T, error)
	Restore(c *gin.Context, id uint) (bool, error)
}

func (s *Service[T]) extended() (extendedRepo[T], error) {
	r, ok := (*s.Repo).(extendedRepo[T])
	if !ok {
		return nil, errors.New(`repo 不支持扩展操作（事务/批量/恢复）`)
	}
	return r, nil
}

// Transaction 在事务中执行 fn。
func (s *Service[T]) Transaction(fn func(tx *gorm.DB) error) error {
	r, err := s.extended()
	if err != nil {
		return err
	}
	return r.Transaction(fn)
}

// BatchAdd 批量新增。
func (s *Service[T]) BatchAdd(c *gin.Context, models []T) error {
	r, err := s.extended()
	if err != nil {
		return err
	}
	return r.BatchAdd(c, models)
}

// BatchDelete 批量软删除。
func (s *Service[T]) BatchDelete(c *gin.Context, ids []uint) error {
	r, err := s.extended()
	if err != nil {
		return err
	}
	return r.BatchDelete(c, ids)
}

// UnscopedOne 查询一条记录（包含软删除）。
func (s *Service[T]) UnscopedOne(c *gin.Context, id uint) (T, error) {
	r, err := s.extended()
	if err != nil {
		var zero T
		return zero, err
	}
	return r.UnscopedOne(c, id)
}

// Restore 恢复软删除的记录。
func (s *Service[T]) Restore(c *gin.Context, id uint) (bool, error) {
	r, err := s.extended()
	if err != nil {
		return false, err
	}
	return r.Restore(c, id)
}
