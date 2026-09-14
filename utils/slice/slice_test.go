package slice

import (
	"reflect"
	"testing"
)

func TestDifferenceAndDirrerenceAlias(t *testing.T) {
	a := []string{"a", "b", "c"}
	b := []string{"b"}

	got := Difference(a, b)
	want := []string{"a", "c"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Difference(%v, %v) = %v, want %v", a, b, got, want)
	}

	// Dirrerence 是 Difference 的历史别名，结果应一致
	if alias := Dirrerence(a, b); !reflect.DeepEqual(alias, want) {
		t.Fatalf("Dirrerence(%v, %v) = %v, want %v", a, b, alias, want)
	}
}

func TestIntersect(t *testing.T) {
	got := Intersect([]string{"a", "b", "c"}, []string{"b", "c", "d"})
	want := []string{"b", "c"}
	if len(got) != len(want) {
		t.Fatalf("Intersect = %v, want %v", got, want)
	}
	for _, w := range want {
		if _, ok := contains(got, w); !ok {
			t.Fatalf("Intersect missing %q: %v", w, got)
		}
	}
}

func TestUnion(t *testing.T) {
	got := Union([]string{"a", "b"}, []string{"b", "c"})
	if len(got) != 3 {
		t.Fatalf("Union = %v, want 3 elements", got)
	}
}

func TestChunk(t *testing.T) {
	got := Chunk([]int{1, 2, 3, 4, 5}, 2)
	want := [][]int{{1, 2}, {3, 4}, {5}}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Chunk = %v, want %v", got, want)
	}
}

func contains(s []string, v string) (int, bool) {
	for i, x := range s {
		if x == v {
			return i, true
		}
	}
	return -1, false
}

func TestIsInArray(t *testing.T) {
	if exists, idx := IsInArray("b", []string{"a", "b", "c"}); !exists || idx != 1 {
		t.Errorf("IsInArray found case: exists=%v idx=%d", exists, idx)
	}
	if exists, idx := IsInArray("z", []string{"a", "b"}); exists || idx != -1 {
		t.Errorf("IsInArray not-found case: exists=%v idx=%d", exists, idx)
	}
	if exists, idx := IsInArray(2, []int{1, 2, 3}); !exists || idx != 1 {
		t.Errorf("IsInArray int-slice case: exists=%v idx=%d", exists, idx)
	}
	// 非 slice 类型返回 (false, -1)
	if exists, idx := IsInArray(1, "not-a-slice"); exists || idx != -1 {
		t.Errorf("IsInArray non-slice case: exists=%v idx=%d", exists, idx)
	}
}

func TestRemoveEmptyString(t *testing.T) {
	got := RemoveEmptyString([]string{" a ", "", "  ", "b", "c"})
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("RemoveEmptyString = %v, want %v", got, want)
	}
}
