/*
Copyright © 2023 jun<simpleyuan@gmail.com>
*/
package comer

import "runtime"

const version = `v1.3.19`

func (c *Comer) goVersion() string {
	return runtime.Version()
}

// Version 返回版本号。
func (c *Comer) Version() string {
	return version
}
