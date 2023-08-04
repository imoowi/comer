/*
Copyright © 2023 jun<simpleyuan@gmail.com>
*/
package cmd

import (
	"github.com/imoowi/comer/comer"
	"github.com/spf13/cobra"
)

// addCmd represents the generate command
var addCmd = &cobra.Command{
	Use:          "add",
	Short:        "Add an app",
	Example:      "comer add",
	SilenceUsage: true,
	PreRun: func(cmd *cobra.Command, args []string) {
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		comerIns := comer.NewComer()
		comerIns.GenApp(cmd, args)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// addCmd.PersistentFlags().String("app", "", "模块名")
	addCmd.PersistentFlags().StringP(`app`, `a`, ``, `模块名`)
	// addCmd.PersistentFlags().String("swaggerTags", "", "接口文档模块名")
	addCmd.PersistentFlags().StringP(`swaggerTags`, `w`, ``, `接口文档模块名`)
	// addCmd.PersistentFlags().String("handler", "", "处理器名")
	addCmd.PersistentFlags().StringP(`handler`, `c`, ``, `处理器|控制器名`)
	// addCmd.PersistentFlags().String("service", "", "服务名")
	addCmd.PersistentFlags().StringP(`service`, `s`, ``, `服务名`)
	// addCmd.PersistentFlags().String("model", "", "模型名;多个模型名之间用英文半角逗号(,)分隔")
	addCmd.PersistentFlags().StringP(`model`, `m`, ``, `模型名;多个模型名之间用英文半角逗号(,)分隔`)
}
