/*
Copyright © 2023 jun<simpleyuan@gmail.com>
*/
package cmd

import (
	"github.com/imoowi/comer/comer"
	"github.com/spf13/cobra"
)

// removeCmd 删除一个应用/控制器（v2 布局）
var removeCmd = &cobra.Command{
	Use:          "remove",
	Short:        "Remove an app (v2 layout)",
	Example:      "comer remove -c=student",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return comer.NewComer().RemoveApp(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)

	removeCmd.PersistentFlags().StringP(`controller`, `c`, ``, `处理器|控制器名`)
	removeCmd.PersistentFlags().StringP(`service`, `s`, ``, `服务名`)
	removeCmd.PersistentFlags().StringP(`model`, `m`, ``, `模型名`)
	removeCmd.PersistentFlags().Bool(`dry-run`, false, `只打印将删除的文件，不实际删除`)
}
