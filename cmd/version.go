/*
Copyright © 2023 jun<simpleyuan@gmail.com>
*/
package cmd

import (
	"github.com/imoowi/comer/comer"
	"github.com/spf13/cobra"
)

// versionCmd 打印版本号
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "version",
	Long:  `version`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("Comer version ", comer.NewComer().Version())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
