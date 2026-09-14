package components

import (
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
)

func TestRedisCaptchaStore(t *testing.T) {
	mr := miniredis.RunT(t)
	client := redis.NewClient(&redis.Options{Addr: mr.Addr()})
	defer client.Close()

	s := NewRedisCaptchaStore(client, time.Minute, "captcha:")

	// Set + Get
	if err := s.Set("id1", "AbCd"); err != nil {
		t.Fatal(err)
	}
	if got := s.Get("id1", false); got != "AbCd" {
		t.Errorf("Get = %q, want AbCd", got)
	}
	// Get with clear 会删除
	if got := s.Get("id1", true); got != "AbCd" {
		t.Errorf("Get(clear) = %q, want AbCd", got)
	}
	if got := s.Get("id1", false); got != "" {
		t.Errorf("after clear Get = %q, want empty", got)
	}
	// Verify 大小写不敏感，且 verify 后清除
	_ = s.Set("id2", "hello")
	if !s.Verify("id2", "HELLO", true) {
		t.Errorf("Verify should be case-insensitive")
	}
	if s.Verify("id2", "HELLO", true) {
		t.Errorf("Verify after clear should be false")
	}
	// 不存在返回空
	if got := s.Get("nope", false); got != "" {
		t.Errorf("Get nonexistent = %q, want empty", got)
	}
}
