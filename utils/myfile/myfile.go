/*
Copyright © 2023 jun<simpleyuan@gmail.com>
*/
package myfile

import (
	"os"
)

func IsFileExist(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)

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
		return err == nil
	}
	return false
}

// 删除目录
func RemoveDir(dirName string) error {
	_err := os.RemoveAll(dirName)
	return _err
}
