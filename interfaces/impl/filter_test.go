package impl

import "testing"

func TestBuildPageListFilterClamp(t *testing.T) {
	f := &Filter{}
	f.SetPage(0)
	f.SetPageSize(0)
	f.BuildPageListFilter(nil, nil)
	if f.GetPage() != 1 || f.GetPageSize() != 20 {
		t.Errorf("clamp defaults: page=%d pageSize=%d, want 1/20", f.GetPage(), f.GetPageSize())
	}

	f.SetPage(5)
	f.SetPageSize(5000)
	f.BuildPageListFilter(nil, nil)
	if f.GetPage() != 5 || f.GetPageSize() != 1000 {
		t.Errorf("clamp max: page=%d pageSize=%d, want 5/1000", f.GetPage(), f.GetPageSize())
	}

	f.SetPage(3)
	f.SetPageSize(50)
	f.BuildPageListFilter(nil, nil)
	if f.GetPage() != 3 || f.GetPageSize() != 50 {
		t.Errorf("valid range unchanged: page=%d pageSize=%d", f.GetPage(), f.GetPageSize())
	}
}
