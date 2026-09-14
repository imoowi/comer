package response

import "testing"

func TestMakePages(t *testing.T) {
	p := MakePages(100, 1, 20)
	if p.Count != 100 || p.CurPage != 1 || p.PageSize != 20 || p.TotalPage != 5 {
		t.Errorf("MakePages(100,1,20) = %+v", p)
	}
	// 不能整除 → 向上取整
	p = MakePages(101, 1, 20)
	if p.TotalPage != 6 {
		t.Errorf("MakePages(101,1,20).TotalPage = %d, want 6", p.TotalPage)
	}
	// pageSize == 0 → 除零保护，TotalPage 为 0
	p = MakePages(100, 1, 0)
	if p.TotalPage != 0 {
		t.Errorf("MakePages(100,1,0).TotalPage = %d, want 0", p.TotalPage)
	}
	// count == 0
	p = MakePages(0, 1, 20)
	if p.TotalPage != 0 || p.Count != 0 {
		t.Errorf("MakePages(0,1,20) = %+v", p)
	}
}
