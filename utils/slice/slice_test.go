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
