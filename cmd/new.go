/*
Copyright © 2023 jun<simpleyuan@gmail.com>
*/
package cmd

import (
	"github.com/imoowi/comer/comer"
	"github.com/spf13/cobra"
)

// createCmd 新建项目
var createCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a project",
	Long:  `comer new [module]`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return comer.NewComer().Start(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.PersistentFlags().String(`config`, ``, `配置文件（JSON5，可覆盖 db_name/exe_name/swagger）`)
}
