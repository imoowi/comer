package comer

import "testing"

func TestComer(t *testing.T) {
	comerIns := NewComer()
	if comerIns.Version() == `` {
		t.Errorf(`comer.Version() test failed`)
	}
}
