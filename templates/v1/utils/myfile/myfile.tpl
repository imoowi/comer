/*
Copyright © 2023 yuanjun<simpleyuan@gmail.com>
*/
package myfile

import (
	"fmt"
	"os"
)

func IsFileExist(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		// fmt.Println(info)
		return false
	}
	fmt.Println("exists", info.Name(), info.Size(), info.ModTime())
	return true
}

// 创建目录
func CreateDir(dirName string, remove bool) bool {
	createFlag := false
	if !IsFileExist(dirName) {
		createFlag = true
	} else {
		if remove {
			RemoveDir(dirName)
			createFlag = true
		}
	}

	if createFlag {
		err := os.MkdirAll(dirName, 0755)
		if err != nil {
			return false
		}
		return true
	}
	return false
}

// 删除目录
func RemoveDir(dirName string) error {
	_err := os.RemoveAll(dirName)
	return _err
}
