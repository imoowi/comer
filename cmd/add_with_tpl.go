/*
Copyright © 2023 jun<simpleyuan@gmail.com>
*/
package cmd

import (
	"github.com/imoowi/comer/comer"
	"github.com/spf13/cobra"
)

// addWithTplCmd 通过自定义模板添加应用
var addWithTplCmd = &cobra.Command{
	Use:          "add-with-tpl",
	Short:        "Add an app with templates",
	Example:      "comer add-with-tpl -t=.comer-templates",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return comer.NewComer().GenAppWithTpl(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(addWithTplCmd)
	addWithTplCmd.PersistentFlags().StringP(`tpl`, `t`, `.comer-templates`, `模板目录`)
}
