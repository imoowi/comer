/*
Copyright © 2023 jun<simpleyuan@gmail.com>
*/
package comer

import (
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
	path       string
	App        *App
	tplAppData map[string]any
}

type Framework struct {
	dirs  []string
	files map[string]string
}
type App struct {
	dirs  []string
	files map[string]string
}

func (c *Comer) Start(cmd *cobra.Command, args []string) {
	if !c.init(cmd, args) {
		return
	}
	c.showLogo()
	c.generateFrameworkDir()
	c.generateFrameworkFiles()
	c.showTips()
}

func (c *Comer) generateFrameworkDir() {

	if len(c.Framework.dirs) > 0 {
		for _, dir := range c.Framework.dirs {
			c.generateDirByName(dir)
		}
	}
}

func (c *Comer) generateFrameworkFiles() {
	if len(c.Framework.files) > 0 {
		for file, tpl := range c.Framework.files {
			c.generateFileByMap(file, tpl, c.tplData)
		}
	}
}
