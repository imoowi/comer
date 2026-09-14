/*
Copyright © 2023 jun<simpleyuan@gmail.com>
*/
package components

import (
	"context"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
)

// RedisCaptchaStore 基于 Redis 的验证码存储，实现 base64Captcha.Store。
type RedisCaptchaStore struct {
	client redis.Cmdable
	ttl    time.Duration
	prefix string
}

// NewRedisCaptchaStore 创建 Redis 验证码存储。
func NewRedisCaptchaStore(client redis.Cmdable, ttl time.Duration, prefix string) *RedisCaptchaStore {
	if ttl <= 0 {
		ttl = 10 * time.Minute
	}
	if prefix == `` {
		prefix = `captcha:`
	}
	return &RedisCaptchaStore{client: client, ttl: ttl, prefix: prefix}
}

func (s *RedisCaptchaStore) key(id string) string {
	return s.prefix + id
}

func (s *RedisCaptchaStore) Set(id string, value string) error {
	return s.client.Set(context.Background(), s.key(id), value, s.ttl).Err()
}

func (s *RedisCaptchaStore) Get(id string, clear bool) string {
	ctx := context.Background()
	key := s.key(id)
	val, err := s.client.Get(ctx, key).Result()
	if err != nil {
		return ``
	}
	if clear {
		_ = s.client.Del(ctx, key).Err()
	}
	return val
}

func (s *RedisCaptchaStore) Verify(id, answer string, clear bool) bool {
	val := s.Get(id, clear)
	return val != `` && strings.EqualFold(val, answer)
}
