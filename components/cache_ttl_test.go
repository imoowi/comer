package components

import (
	"testing"
	"time"
)

func TestMemCacheTTL(t *testing.T) {
	c := NewMemCacheT[string]()
	c.SetOneTTL("ttl-key", "v", 30*time.Millisecond)
	if v, ok := c.GetOne("ttl-key"); !ok || v != "v" {
		t.Fatalf("expected value before expiry, got %q ok=%v", v, ok)
	}
	time.Sleep(50 * time.Millisecond)
	if _, ok := c.GetOne("ttl-key"); ok {
		t.Fatalf("expected expired key to be missing")
	}
}

func TestMemCacheFlush(t *testing.T) {
	c1 := NewMemCacheT[string]()
	c2 := NewMemCacheT[int]()
	c1.SetOne("a", "1")
	c1.SetOne("b", "2")
	c2.SetOne("a", 99)

	c1.Flush()
	if _, ok := c1.GetOne("a"); ok {
		t.Fatalf("c1 key 'a' should be flushed")
	}
	if _, ok := c1.GetOne("b"); ok {
		t.Fatalf("c1 key 'b' should be flushed")
	}
	if v, ok := c2.GetOne("a"); !ok || v != 99 {
		t.Fatalf("c2 key 'a' should remain, got %v ok=%v", v, ok)
	}
}

func TestFlushAll(t *testing.T) {
	c := NewMemCacheT[string]()
	c.SetOne("flushall-key", "1")
	FlushAll()
	if _, ok := c.GetOne("flushall-key"); ok {
		t.Fatalf("FlushAll should clear everything")
	}
}
