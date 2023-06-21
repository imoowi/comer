package comer

import (
	"fmt"
	"runtime"
)

func (c *Comer) goVersion() string {
	return runtime.Version()
}

func (c *Comer) Version() string {
	c.version = `v1.1.2`
	fmt.Println(`Comer version `, c.version)
	return c.version
}
