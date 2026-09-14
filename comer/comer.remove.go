/*
Copyright © 2023 jun<simpleyuan@gmail.com>
*/
package comer

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// RemoveApp 删除 add 为某 controller 生成的 v2 文件（add 的逆操作）。
func (c *Comer) RemoveApp(cmd *cobra.Command, args []string) error {
	tplVersion, err := cmd.Flags().GetString(`tplVersion`)
	if err != nil {
		return err
	}
	if tplVersion == `1` {
		return fmt.Errorf(`remove 仅支持 v2 布局`)
	}

	controllerName, err := cmd.Flags().GetString(`controller`)
	if err != nil {
		return err
	}
	if controllerName == `` {
		return fmt.Errorf(`pls input controller name, e.g. -c=student (请输入控制器名,例如 -c=student)`)
	}
	serviceName, _ := cmd.Flags().GetString(`service`)
	if serviceName == `` {
		serviceName = controllerName
	}
	modelName, _ := cmd.Flags().GetString(`model`)
	if modelName == `` {
		modelName = serviceName
	}
	dryRun, _ := cmd.Flags().GetBool(`dry-run`)

	// 保护：避免在非项目根目录误删
	if _, err := os.Stat(`go.mod`); os.IsNotExist(err) {
		return fmt.Errorf(`当前目录下没有 go.mod，请在项目根目录执行 remove`)
	}

	removed := 0
	for p := range appFilePathsV2(controllerName, serviceName, modelName) {
		if _, err := os.Stat(p); os.IsNotExist(err) {
			if dryRun {
				fmt.Printf("missing %s\n", p)
			}
			continue
		}
		if dryRun {
			fmt.Printf("would remove %s\n", p)
			continue
		}
		if err := os.Remove(p); err != nil {
			return err
		}
		fmt.Printf("removed %s\n", p)
		removed++
	}

	if !dryRun && removed == 0 {
		fmt.Println(`nothing to remove`)
	}
	return nil
}
