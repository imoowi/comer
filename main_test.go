package main

import "testing"

func TestMain(t *testing.T) {
	if 1 == 2 {
		t.Errorf(`main test failed`)
	}
}
