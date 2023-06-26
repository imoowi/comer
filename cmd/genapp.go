/*
Copyright © 2023 yuanjun<simpleyuan@gmail.com>
*/
package cmd

import (
	"github.com/imoowi/comer/comer"
	"github.com/spf13/cobra"
)

// genappCmd represents the generate command
var genappCmd = &cobra.Command{
	Use:          "genapp",
	Short:        "app生成器",
	Example:      "comer genapp",
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
	rootCmd.AddCommand(genappCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// genappCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// genappCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	genappCmd.PersistentFlags().String("app", "", "模块名")
	genappCmd.PersistentFlags().String("swaggerTags", "", "接口文档模块名")
	genappCmd.PersistentFlags().String("handler", "", "处理器名")
	genappCmd.PersistentFlags().String("service", "", "服务名")
	genappCmd.PersistentFlags().String("model", "", "模型名;多个模型名之间用英文半角逗号(,)分隔")
}
