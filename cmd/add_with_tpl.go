/*
Copyright © 2023 jun<simpleyuan@gmail.com>
*/
package cmd

import (
	"github.com/imoowi/comer/comer"
	"github.com/spf13/cobra"
)

// addWithTplCmd represents the generate command
var addWithTplCmd = &cobra.Command{
	Use:          "add-with-tpl",
	Short:        "Add an app with templates",
	Example:      "comer add-with-tpl -t=.comer-templates",
	SilenceUsage: true,
	PreRun: func(cmd *cobra.Command, args []string) {
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		comerIns := comer.NewComer()
		comerIns.GenAppWithTpl(cmd, args)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addWithTplCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addWithTplCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addWithTplCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// addWithTplCmd.PersistentFlags().String("app", "", "模块名")
	addWithTplCmd.PersistentFlags().StringP(`tpl`, `t`, `.comer-templates`, `模板目录`)
}
