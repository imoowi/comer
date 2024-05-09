package components

import "sync"

var MemCache sync.Map

type MemCacheT[T any] struct {
}

func NewMemCacheT[T any]() *MemCacheT[T] {
	return &MemCacheT[T]{}
}
func (c *MemCacheT[T]) GetOne(key string) (val T, ok bool) {
	v, ok := c.load(key)
	if v != nil {
		val = v.(T)
	}
	return
}
func (c *MemCacheT[T]) SetOne(key string, value T) {
	c.store(key, value)
}

func (c *MemCacheT[T]) GetArray(key string) (val []T, ok bool) {
	v, ok := c.load(key)
	if v != nil {
		val = v.([]T)
	}
	return
}
func (c *MemCacheT[T]) SetArray(key string, values []T) {
	c.store(key, values)
}

func (c *MemCacheT[T]) Del(key string) {
	MemCache.Delete(key)
}

func (c *MemCacheT[T]) load(key string) (val any, ok bool) {
	if v, ok := MemCache.Load(key); ok {
		val = v
		return val, ok
	}
	return
}
func (c *MemCacheT[T]) store(key string, values any) {
	MemCache.Store(key, values)
}
