package comer

import "testing"

func TestComer(t *testing.T) {
	comerIns := NewComer()
	if comerIns.Version() == `` {
		t.Errorf(`comer.Version() test failed`)
	}
	if comerIns.goVersion() == `` {
		t.Errorf(`comer.goVersion() test failed`)
	}
	if comerIns.showTips() != `showTips called` {
		t.Errorf(`showTips test failed`)
	}
}
