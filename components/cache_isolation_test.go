package components

import "testing"

func TestMemCacheTypeIsolation(t *testing.T) {
	c1 := NewMemCacheT[string]()
	c2 := NewMemCacheT[int]()

	c1.SetOne("isolated-key", "hello")
	if _, ok := c2.GetOne("isolated-key"); ok {
		t.Fatal("cross-type key collision: c2[int] should not see c1[string]'s key")
	}
}

func TestMemCacheTypeMismatchNoPanic(t *testing.T) {
	c := NewMemCacheT[string]()
	c.SetOne("mismatch-key", "a string")

	// 用错误的方法读取不同类型，应返回 ok=false 而非 panic
	if _, ok := c.GetArray("mismatch-key"); ok {
		t.Fatal("expected ok=false when reading a string as []string")
	}
}
