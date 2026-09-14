package impl

import (
	"testing"
	"time"

	"github.com/glebarez/sqlite"
	"github.com/imoowi/comer/interfaces"
	"gorm.io/gorm"
)

type testModel struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string
}

func (m *testModel) GetID() uint       { return m.ID }
func (m *testModel) SetId(id uint)     { m.ID = id }
func (m *testModel) TableName() string { return "test_models" }

func newTestRepo(t *testing.T) *Repo[*testModel] {
	t.Helper()
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	if err := db.AutoMigrate(&testModel{}); err != nil {
		t.Fatal(err)
	}
	return NewRepo[*testModel](db)
}

func newFilter(page, pageSize int64) *interfaces.IFilter {
	f := &Filter{}
	f.SetPage(page)
	f.SetPageSize(pageSize)
	var iface interfaces.IFilter = f
	return &iface
}

func add(t *testing.T, r *Repo[*testModel], name string) uint {
	t.Helper()
	id, err := r.Add(nil, &testModel{Name: name})
	if err != nil {
		t.Fatal(err)
	}
	return id
}

func TestRepoAdd(t *testing.T) {
	r := newTestRepo(t)
	if id := add(t, r, "foo"); id == 0 {
		t.Fatal("Add should return non-zero id")
	}
}

func TestRepoOne(t *testing.T) {
	r := newTestRepo(t)
	id := add(t, r, "foo")
	got, err := r.One(nil, id)
	if err != nil {
		t.Fatal(err)
	}
	if got.Name != "foo" {
		t.Errorf("One = %+v, want Name foo", got)
	}
}

func TestRepoOneWithSelectOption(t *testing.T) {
	r := newTestRepo(t)
	id := add(t, r, "foo")
	got, err := r.OneWithSelectOption(nil, id, []string{"name"})
	if err != nil {
		t.Fatal(err)
	}
	if got.Name != "foo" {
		t.Errorf("OneWithSelectOption = %+v", got)
	}
}

func TestRepoOneByName(t *testing.T) {
	r := newTestRepo(t)
	add(t, r, "foo")
	got, err := r.OneByName(nil, "foo")
	if err != nil {
		t.Fatal(err)
	}
	if got.Name != "foo" {
		t.Errorf("OneByName = %+v", got)
	}
}

func TestRepoPageList(t *testing.T) {
	r := newTestRepo(t)
	for _, n := range []string{"a", "b", "c"} {
		add(t, r, n)
	}
	res, err := r.PageList(nil, newFilter(1, 10))
	if err != nil {
		t.Fatal(err)
	}
	if len(res.List) != 3 || res.Pages.Count != 3 {
		t.Errorf("PageList = %+v", res)
	}
}

func TestRepoUpdate(t *testing.T) {
	r := newTestRepo(t)
	id := add(t, r, "foo")
	updated, err := r.Update(nil, map[string]any{"name": "bar"}, id)
	if err != nil || !updated {
		t.Fatalf("Update = %v, %v", updated, err)
	}
	got, _ := r.One(nil, id)
	if got.Name != "bar" {
		t.Errorf("after update Name = %q, want bar", got.Name)
	}
}

func TestRepoDelete(t *testing.T) {
	r := newTestRepo(t)
	id := add(t, r, "foo")
	deleted, err := r.Delete(nil, id)
	if err != nil || !deleted {
		t.Fatalf("Delete = %v, %v", deleted, err)
	}
	if _, err := r.One(nil, id); err == nil {
		t.Fatal("One after Delete should return record-not-found")
	}
}

func TestRepoTransaction(t *testing.T) {
	r := newTestRepo(t)
	err := r.Transaction(func(tx *gorm.DB) error {
		return tx.Create(&testModel{Name: "tx"}).Error
	})
	if err != nil {
		t.Fatal(err)
	}
	if _, err := r.OneByName(nil, "tx"); err != nil {
		t.Fatalf("transaction-inserted row not found: %v", err)
	}
}

func TestRepoBatchAdd(t *testing.T) {
	r := newTestRepo(t)
	models := []*testModel{{Name: "a"}, {Name: "b"}}
	if err := r.BatchAdd(nil, models); err != nil {
		t.Fatal(err)
	}
	if res, _ := r.PageList(nil, newFilter(1, 10)); len(res.List) != 2 {
		t.Errorf("BatchAdd inserted %d, want 2", len(res.List))
	}
}

func TestRepoBatchDelete(t *testing.T) {
	r := newTestRepo(t)
	id1 := add(t, r, "a")
	id2 := add(t, r, "b")
	if err := r.BatchDelete(nil, []uint{id1, id2}); err != nil {
		t.Fatal(err)
	}
	if res, _ := r.PageList(nil, newFilter(1, 10)); len(res.List) != 0 {
		t.Errorf("BatchDelete left %d rows", len(res.List))
	}
}

func TestRepoUnscopedOneAndRestore(t *testing.T) {
	r := newTestRepo(t)
	id := add(t, r, "foo")
	if _, err := r.Delete(nil, id); err != nil {
		t.Fatal(err)
	}
	if _, err := r.UnscopedOne(nil, id); err != nil {
		t.Fatalf("UnscopedOne: %v", err)
	}
	restored, err := r.Restore(nil, id)
	if err != nil || !restored {
		t.Fatalf("Restore = %v, %v", restored, err)
	}
	if _, err := r.One(nil, id); err != nil {
		t.Fatalf("One after Restore should find row: %v", err)
	}
}

func TestValidateID(t *testing.T) {
	if err := validateID(0); err == nil {
		t.Fatal("validateID(0) should error")
	}
	if err := validateID(1); err != nil {
		t.Fatal("validateID(1) should not error")
	}
}
