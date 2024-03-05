/*
Copyright © 2023 jun<simpleyuan@gmail.com>
*/
package comer

import (
	"fmt"
	"runtime"
)

func (c *Comer) goVersion() string {
	return runtime.Version()
}

func (c *Comer) Version() string {
	c.version = `v1.3.2`
	fmt.Println(`Comer version `, c.version)
	return c.version
}
