package slice

import (
	"reflect"
	"strings"
)

// 并集
func Union(slice1, slice2 []string) []string {
	m := make(map[string]int)
	for _, v := range slice1 {
		m[v]++
	}

	for _, v := range slice2 {
		times, _ := m[v]
		if times == 0 {
			slice1 = append(slice1, v)
		}
	}
	return slice1
}

// 交集
func Intersect(slice1, slice2 []string) []string {
	m := make(map[string]int)
	nn := make([]string, 0)
	for _, v := range slice1 {
		m[v]++
	}

	for _, v := range slice2 {
		times, _ := m[v]
		if times == 1 {
			nn = append(nn, v)
		}
	}
	return nn
}

// 差集
func Dirrerence(slice1, slice2 []string) []string {
	m := make(map[string]int)
	nn := make([]string, 0)
	inter := Intersect(slice1, slice2)
	for _, v := range inter {
		m[v]++
	}

	for _, value := range slice1 {
		times, _ := m[value]
		if times == 0 {
			nn = append(nn, value)
		}
	}
	return nn
}

func IsInArray(val interface{}, array interface{}) (exists bool, index int) {
	exists = false
	index = -1

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				index = i
				exists = true
				return
			}
		}
	}

	return
}

// 从切片中去除空字符串
func RemoveEmptyString(source []string) []string {
	res := []string{}
	for _, str := range source {
		str := strings.TrimSpace(str)
		if str != "" {
			res = append(res, str)
		}
	}
	return res
}

// 对切片进行分隔
func Chunk[T any](src []T, size int) [][]T {
	chunked := make([][]T, 0)
	var chunkItem []T
	resIdx := -1
	for idx, item := range src {
		m := idx % size
		if m == 0 {
			resIdx++
			chunkItem = make([]T, 0)
			chunked = append(chunked, chunkItem)
		}
		chunkItem = append(chunkItem, item)
		chunked[resIdx] = chunkItem
	}

	return chunked
}
