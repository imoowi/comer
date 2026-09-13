/*
Copyright © 2023 jun<simpleyuan@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd 是根命令
var rootCmd = &cobra.Command{
	Use:   "comer",
	Short: "comer",
	Long:  `comer`,
}

// Execute 由 main.main() 调用，执行根命令。
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringP(`tplVersion`, `v`, ``, `模板版本号`)
}
