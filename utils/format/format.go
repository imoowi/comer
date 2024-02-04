/*
Copyright © 2023 jun<simpleyuan@gmail.com>
*/
package format

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func FirstUpper(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

func FirstLower(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToLower(s[:1]) + s[1:]
}

// 驼峰转短横线
func Camel2Dash(s string) string {
	buf := make([]rune, 0)
	for i, v := range s {
		if unicode.IsUpper(v) {
			if i != 0 {
				buf = append(buf, []rune(`-`)[0])
			}
			buf = append(buf, unicode.ToLower(v))
		} else {
			buf = append(buf, v)
		}
	}
	return string(buf)
}

// 驼峰转下划线
func Camel2Snake(s string) string {
	buf := make([]rune, 0)
	for i, v := range s {
		if unicode.IsUpper(v) {
			if i != 0 {
				buf = append(buf, []rune(`_`)[0])
			}
			buf = append(buf, unicode.ToLower(v))
		} else {
			buf = append(buf, v)
		}
	}
	return string(buf)
}

// 16进制字符串转10进制
func Hex2Dec(val string) int {
	val = strings.Replace(val, `0x`, ``, -1)
	val = strings.Replace(val, `0X`, ``, -1)
	n, err := strconv.ParseUint(val, 16, 32)
	if err != nil {
		fmt.Println(err)
	}
	return int(n)
}

// 对一个string切片进行去重
func UniqueSliceString(input []string) []string {
	strMap := make(map[string]bool)
	for _, str := range input {
		strMap[str] = true
	}
	res := make([]string, len(strMap))
	idx := 0
	for str := range strMap {
		res[idx] = str
		idx++
	}
	return res
}
