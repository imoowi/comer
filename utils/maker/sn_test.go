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
