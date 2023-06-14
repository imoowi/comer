package comer

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

type Comer struct {
	cfgPath     string
	debugMode   bool
	version     string
	showVersion bool
	cmd         *cobra.Command
	args        []string
}

func (c *Comer) Start(cmd *cobra.Command, args []string) {
	c.init(cmd, args)
	fmt.Printf(`
	_________                                   
	\_   ___ \   ____    _____    ____  _______ 
	/    \  \/  /  _ \  /     \ _/ __ \ \_  __ \
	\     \____(  <_> )|  Y Y  \\  ___/  |  | \/
	 \______  / \____/ |__|_|  / \___  > |__|   
			\/               \/      \/ %s, built with %s

`, c.Version(), c.goVersion())

	if c.showVersion {
		return
	}

	if c.debugMode {
		fmt.Println("[debug] mode")
	}
}

func (c *Comer) init(cmd *cobra.Command, args []string) {
	c.debugMode = true
	c.cmd = cmd
	c.args = args
	c.version = `v0.1.0`
}

func (c *Comer) Version() string {
	c.init(nil, nil)
	fmt.Println(`comer version `, c.version)
	return c.version
}
func (c *Comer) goVersion() string {
	return runtime.Version()
}
func NewComer() *Comer {
	return &Comer{}
}
