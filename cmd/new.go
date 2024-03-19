/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/imoowi/comer/comer"
	"github.com/spf13/cobra"
)

// createCmd represents the new command
var createCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a project",
	Long:  `comer new [module]`,
	Run: func(cmd *cobra.Command, args []string) {
		comerIns := comer.NewComer()
		comerIns.Start(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// createCmd.PersistentFlags().StringP(`module`, `m`, ``, `go.mod module(go.mod文件的module名称)`)
	// createCmd.MarkFlagRequired(`module`)
	// createCmd.PersistentFlags().String(`path`, ``, `project root (项目所在目录)`)

	rootCmd.PersistentFlags().StringP(`tplVersion`, `v`, ``, `模板版本号`)
}
