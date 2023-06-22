package cmd

import "testing"

func TestCmd(t *testing.T) {
	if 1 == 2 {
		t.Errorf(`main test failed`)
	}
}
