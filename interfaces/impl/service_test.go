package impl

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/imoowi/comer/interfaces"
	"github.com/imoowi/comer/utils/response"
	"gorm.io/gorm"
)

// mockRepo 只实现 IRepo[T]，不实现 extendedRepo[T]，用于测 Service 扩展方法的类型断言失败分支。
type mockRepo[T interfaces.IModel] struct{}

func (m *mockRepo[T]) PageList(c *gin.Context, f *interfaces.IFilter) (*response.PageListT[T], error) {
	return nil, nil
}
func (m *mockRepo[T]) PageListWithSelectOption(c *gin.Context, f *interfaces.IFilter, s []string) (*response.PageListT[T], error) {
	return nil, nil
}
func (m *mockRepo[T]) One(c *gin.Context, id uint) (T, error) {
	var z T
	return z, nil
}
func (m *mockRepo[T]) OneWithSelectOption(c *gin.Context, id uint, s []string) (T, error) {
	var z T
	return z, nil
}
func (m *mockRepo[T]) OneByName(c *gin.Context, name string) (T, error) {
	var z T
	return z, nil
}
func (m *mockRepo[T]) OneByNameWithSelectOption(c *gin.Context, name string, s []string) (T, error) {
	var z T
	return z, nil
}
func (m *mockRepo[T]) Add(c *gin.Context, model T) (uint, error) { return 0, nil }
func (m *mockRepo[T]) Update(c *gin.Context, u map[string]any, id uint) (bool, error) {
	return false, nil
}
func (m *mockRepo[T]) Delete(c *gin.Context, id uint) (bool, error) { return false, nil }

func TestServiceCRUD(t *testing.T) {
	r := newTestRepo(t)
	svc := NewService[*testModel](r)

	id, err := svc.Add(nil, &testModel{Name: "foo"})
	if err != nil || id == 0 {
		t.Fatalf("svc.Add = %d, %v", id, err)
	}
	got, err := svc.One(nil, id)
	if err != nil || got.Name != "foo" {
		t.Fatalf("svc.One = %+v, %v", got, err)
	}
	// Update 会剥离 created_at/updated_at/deleted_at
	updated, err := svc.Update(nil, map[string]any{"name": "bar", "created_at": "x", "updated_at": "y", "deleted_at": "z"}, id)
	if err != nil || !updated {
		t.Fatalf("svc.Update = %v, %v", updated, err)
	}
	deleted, err := svc.Delete(nil, id)
	if err != nil || !deleted {
		t.Fatalf("svc.Delete = %v, %v", deleted, err)
	}
}

func TestServiceExtendedMethods(t *testing.T) {
	r := newTestRepo(t)
	svc := NewService[*testModel](r)

	if err := svc.Transaction(func(tx *gorm.DB) error {
		return tx.Create(&testModel{Name: "tx"}).Error
	}); err != nil {
		t.Fatal(err)
	}
	if err := svc.BatchAdd(nil, []*testModel{{Name: "a"}, {Name: "b"}}); err != nil {
		t.Fatal(err)
	}
	if _, err := svc.UnscopedOne(nil, 1); err != nil {
		t.Fatalf("svc.UnscopedOne: %v", err)
	}
}

func TestServiceExtendedErrorForMockRepo(t *testing.T) {
	svc := NewService[*testModel](&mockRepo[*testModel]{})
	if err := svc.Transaction(func(tx *gorm.DB) error { return nil }); err == nil {
		t.Fatal("expected error when repo does not support extended methods")
	}
	if err := svc.BatchAdd(nil, nil); err == nil {
		t.Fatal("expected error when repo does not support extended methods")
	}
}
