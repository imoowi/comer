package format

import (
	"strings"
)

func Format(val interface{}) interface{} {
	return val
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

// 正则表达式转义
func RegexStringTranformer(str string) string {
	//正则匹配出现的特殊字符串
	fbsArr := []string{"\\", "$", "(", ")", "*", "+", ".", "[", "]", "?", "^", "{", "}", "|"}
	for _, ch := range fbsArr {
		if StrContainers := strings.Contains(str, ch); StrContainers {
			str = strings.Replace(str, ch, "\\"+ch, -1)
		}
	}
	return str
}

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
