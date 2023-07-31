/*
Copyright © 2023 jun<simpleyuan@gmail.com>
*/
package cmd

import (
	"github.com/imoowi/comer/comer"
	"github.com/spf13/cobra"
)

// addappCmd represents the generate command
var addappCmd = &cobra.Command{
	Use:          "addapp",
	Short:        "Add an App (添加app)",
	Example:      "comer addapp",
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
	rootCmd.AddCommand(addappCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addappCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addappCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addappCmd.PersistentFlags().String("app", "", "模块名")
	addappCmd.PersistentFlags().String("swaggerTags", "", "接口文档模块名")
	addappCmd.PersistentFlags().String("handler", "", "处理器名")
	addappCmd.PersistentFlags().String("service", "", "服务名")
	addappCmd.PersistentFlags().String("model", "", "模型名;多个模型名之间用英文半角逗号(,)分隔")
}
