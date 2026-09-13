/*
Copyright © 2023 jun<simpleyuan@gmail.com>
*/
package comer

import (
	"github.com/spf13/cobra"
)

type Comer struct {
	Framework  *Framework
	tplData    map[string]any
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

func (c *Comer) Start(cmd *cobra.Command, args []string) error {
	tplVersion, err := cmd.Flags().GetString(`tplVersion`)
	if err != nil {
		return err
	}

	if tplVersion == `1` {
		if err = c.init(cmd, args); err != nil {
			return err
		}
	} else {
		if err = c.initV2(cmd, args); err != nil {
			return err
		}
	}

	c.showLogo()
	if err = c.generateFrameworkDir(); err != nil {
		return err
	}
	if err = c.generateFrameworkFiles(); err != nil {
		return err
	}
	c.showTips()
	return nil
}

func (c *Comer) generateFrameworkDir() error {
	for _, dir := range c.Framework.dirs {
		if err := c.generateDirByName(dir); err != nil {
			return err
		}
	}
	return nil
}

func (c *Comer) generateFrameworkFiles() error {
	for file, tpl := range c.Framework.files {
		if err := c.generateFileByMap(file, tpl, c.tplData, false); err != nil {
			return err
		}
	}
	return nil
}
