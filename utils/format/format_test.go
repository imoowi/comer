/*
Copyright © 2023 jun<simpleyuan@gmail.com>
*/
package format

import (
	"fmt"
	"testing"
)

func TestFirstUpper(t *testing.T) {
	abc := `abc`
	Abc := FirstUpper(abc)
	if Abc != `Abc` {
		t.Errorf(`FirstUpper test failed`)
	}
	_abc := ``
	_Abc := FirstUpper(_abc)
	if _Abc != `` {
		t.Errorf(`Empty str FirstUpper test failed`)
	}
}

func TestFirstLower(t *testing.T) {
	Abc := `Abc`
	abc := FirstLower(Abc)
	if abc != `abc` {
		t.Errorf(`FirstLower test failed`)
	}
	_Abc := ``
	_abc := FirstLower(_Abc)
	if _abc != `` {
		t.Errorf(`Empty str FirstUpper test failed`)
	}
}

func TestCamel2Dash(t *testing.T) {
	s := `userRole`
	d := Camel2Dash(s)
	if d != `user-role` {
		t.Errorf(`test failed`)
	}
	fmt.Println(d)
}

func TestCamel2Snake(t *testing.T) {
	s := `userRole`
	d := Camel2Snake(s)
	if d != `user_role` {
		t.Errorf(`test failed`)
	}
	fmt.Println(d)
}

func TestHex2Dec(t *testing.T) {
	cases := map[string]int{
		"0x10": 16,
		"0Xff": 255,
		"a":    10,
		"10":   16,
	}
	for in, want := range cases {
		if got := Hex2Dec(in); got != want {
			t.Errorf("Hex2Dec(%q) = %d, want %d", in, got, want)
		}
	}
}

func TestUniqueSliceString(t *testing.T) {
	got := UniqueSliceString([]string{"a", "b", "a", "c", "b"})
	if len(got) != 3 {
		t.Fatalf("UniqueSliceString length = %d, want 3: %v", len(got), got)
	}
	// 去重结果顺序不确定，按集合断言
	set := map[string]bool{}
	for _, s := range got {
		set[s] = true
	}
	for _, want := range []string{"a", "b", "c"} {
		if !set[want] {
			t.Errorf("UniqueSliceString missing %q: %v", want, got)
		}
	}
}
