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
