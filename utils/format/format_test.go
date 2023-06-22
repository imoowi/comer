package format

import "testing"

func TestFirstUpper(t *testing.T) {
	abc := `abc`
	Abc := FirstUpper(abc)
	if Abc != `Abc` {
		t.Errorf(`FirstUpper test failed`)
	}
}

func TestFirstLower(t *testing.T) {
	Abc := `Abc`
	abc := FirstLower(Abc)
	if abc != `abc` {
		t.Errorf(`FirstLower test failed`)
	}
}
