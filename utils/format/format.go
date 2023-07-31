/*
Copyright © 2023 jun<simpleyuan@gmail.com>
*/
package format

import (
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
