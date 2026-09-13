package components

import (
	"reflect"
	"sync"
)

var MemCache sync.Map

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
	c.store(key, value)
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
	c.store(key, values)
}

func (c *MemCacheT[T]) Del(key string) {
	MemCache.Delete(c.namespaced(key))
}

func (c *MemCacheT[T]) load(key string) (val any, ok bool) {
	return MemCache.Load(c.namespaced(key))
}

func (c *MemCacheT[T]) store(key string, values any) {
	MemCache.Store(c.namespaced(key), values)
}
