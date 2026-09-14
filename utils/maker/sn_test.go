/*
Copyright © 2023 jun<simpleyuan@gmail.com>
*/
package maker

import "testing"

func TestMakeSn(t *testing.T) {
	sn := MakeSn(`sn_`)
	if sn == `` {
		t.Errorf(`MakeSn test failed`)
	}
}

func TestMakeRandStr(t *testing.T) {
	str := MakeRandStr(10)
	if str == `` {
		t.Errorf(`MakeRandStr test failed`)
	}
}

func TestMakeRandNumber(t *testing.T) {
	str := MakeRandNumber(10)
	if len(str) != 10 {
		t.Errorf("MakeRandNumber(10) length = %d, want 10", len(str))
	}
	for _, c := range str {
		if c < '0' || c > '9' {
			t.Errorf("MakeRandNumber contains non-digit: %q", c)
			break
		}
	}
}
