/*
Copyright © 2023 jun<simpleyuan@gmail.com>
*/
package cmd

import (
	"github.com/imoowi/comer/comer"
	"github.com/spf13/cobra"
)

// addCmd 添加一个应用/控制器
var addCmd = &cobra.Command{
	Use:          "add",
	Short:        "Add an app",
	Example:      "comer add -a=app01 -c=handler01 -w='app01' -s=service01 -m=model01",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return comer.NewComer().AddApp(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.PersistentFlags().StringP(`app`, `a`, ``, `模块名`)
	addCmd.PersistentFlags().StringP(`swaggerTags`, `w`, ``, `接口文档模块名`)
	addCmd.PersistentFlags().StringP(`controller`, `c`, ``, `处理器|控制器名`)
	addCmd.PersistentFlags().StringP(`service`, `s`, ``, `服务名`)
	addCmd.PersistentFlags().StringP(`model`, `m`, ``, `模型名;多个模型名之间用英文半角逗号(,)分隔`)
	addCmd.PersistentFlags().StringArrayP(`field`, `f`, nil, `模型字段 name:type[:size][:comment][:validate]，可重复`)
	addCmd.PersistentFlags().String(`fieldConfig`, ``, `模型字段配置文件，一行一个 name:type[:size][:comment][:validate]`)
	addCmd.PersistentFlags().String(`searchColumn`, ``, `分页搜索字段（默认取第一个 string/text 字段）`)
}
