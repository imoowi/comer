package comer

import (
	"fmt"

	"github.com/spf13/cobra"
)

type Comer struct {
	debugMode  bool
	version    string
	cmd        *cobra.Command
	args       []string
	Framework  *Framework
	tplData    map[string]any
	moduleName string
}

type Framework struct {
	dirs  []string
	files map[string]string
}

func (c *Comer) Start(cmd *cobra.Command, args []string) {
	if !c.init(cmd, args) {
		fmt.Println(`init failed`)
		return
	}

	fmt.Printf(`
_________                                   
\_   ___ \   ____    _____    ____  _______ 
/    \  \/  /  _ \  /     \ _/ __ \ \_  __ \
\     \____(  <_> )|  Y Y  \\  ___/  |  | \/
 \______  / \____/ |__|_|  / \___  > |__|   
		\/               \/      \/ %s, built with %s
`, c.Version(), c.goVersion())

	c.generateFrameworkDir()
	c.generateFrameworkFiles()
}

func (c *Comer) generateFrameworkDir() {

	if len(c.Framework.dirs) > 0 {
		for _, dir := range c.Framework.dirs {
			c.generateFrameworkDirByName(dir)
		}
	}
}

func (c *Comer) generateFrameworkFiles() {
	if len(c.Framework.files) > 0 {
		for file, tpl := range c.Framework.files {
			c.generateFrameworkFileByMap(file, tpl, c.tplData)
		}
	}
}
