package components

import (
	"reflect"
	"sync"
	"time"
)

var MemCache sync.Map

// cacheEntry 缓存项，expiresAt 零值表示永不过期。
type cacheEntry struct {
	value     any
	expiresAt time.Time
}

type MemCacheT[T any] struct {
}

func NewMemCacheT[T any]() *MemCacheT[T] {
	return &MemCacheT[T]{}
}

// namespaced 为每个类型 T 增加命名空间前缀，避免不同类型实例共享同一 key 空间导致碰撞
func (c *MemCacheT[T]) namespaced(key string) string {
	return reflect.TypeOf((*T)(nil)).Elem().String() + `:` + key
}

func (c *MemCacheT[T]) GetOne(key string) (val T, ok bool) {
	v, ok := c.load(key)
	if !ok {
		return
	}
	val, ok = v.(T)
	return
}

func (c *MemCacheT[T]) SetOne(key string, value T) {
	c.store(key, value, time.Time{})
}

// SetOneTTL 设置带过期时间的缓存。
func (c *MemCacheT[T]) SetOneTTL(key string, value T, ttl time.Duration) {
	c.store(key, value, time.Now().Add(ttl))
}

func (c *MemCacheT[T]) GetArray(key string) (val []T, ok bool) {
	v, ok := c.load(key)
	if !ok {
		return
	}
	val, ok = v.([]T)
	return
}

func (c *MemCacheT[T]) SetArray(key string, values []T) {
	c.store(key, values, time.Time{})
}

// SetArrayTTL 设置带过期时间的数组缓存。
func (c *MemCacheT[T]) SetArrayTTL(key string, values []T, ttl time.Duration) {
	c.store(key, values, time.Now().Add(ttl))
}

func (c *MemCacheT[T]) Del(key string) {
	MemCache.Delete(c.namespaced(key))
}

// Flush 清空当前类型命名空间下的所有缓存。
func (c *MemCacheT[T]) Flush() {
	prefix := reflect.TypeOf((*T)(nil)).Elem().String() + `:`
	MemCache.Range(func(k, _ any) bool {
		if s, ok := k.(string); ok && len(s) >= len(prefix) && s[:len(prefix)] == prefix {
			MemCache.Delete(k)
		}
		return true
	})
}

// FlushAll 清空所有类型的缓存。
func FlushAll() {
	MemCache.Range(func(k, _ any) bool {
		MemCache.Delete(k)
		return true
	})
}

func (c *MemCacheT[T]) load(key string) (val any, ok bool) {
	v, ok := MemCache.Load(c.namespaced(key))
	if !ok {
		return nil, false
	}
	e, ok := v.(cacheEntry)
	if !ok {
		// 兼容裸值（理论上不会出现，除非外部直接操作 MemCache）
		return v, true
	}
	if !e.expiresAt.IsZero() && time.Now().After(e.expiresAt) {
		MemCache.Delete(c.namespaced(key))
		return nil, false
	}
	return e.value, true
}

func (c *MemCacheT[T]) store(key string, value any, expiresAt time.Time) {
	MemCache.Store(c.namespaced(key), cacheEntry{value: value, expiresAt: expiresAt})
}
