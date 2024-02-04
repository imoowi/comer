/*
Copyright © 2023 jun<simpleyuan@gmail.com>
*/
package utils

import (
	"path"
	"runtime"
)

func GetCurrentAbPathByExecutable() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath + `../../`
}
