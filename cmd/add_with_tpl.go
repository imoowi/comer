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
	addWithTplCmd.PersistentFlags().StringArrayP(`field`, `f`, nil, `模型字段 name:type[:size][:comment][:validate]，可重复`)
	addWithTplCmd.PersistentFlags().String(`fieldConfig`, ``, `模型字段配置文件，一行一个 name:type[:size][:comment][:validate]`)
	addWithTplCmd.PersistentFlags().String(`searchColumn`, ``, `分页搜索字段（默认取第一个 string/text 字段）`)
	addWithTplCmd.PersistentFlags().Bool(`dry-run`, false, `只打印将生成的目录与文件，不实际生成`)
}
