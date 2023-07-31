/*
Copyright © 2023 jun<simpleyuan@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/imoowi/comer/comer"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "comer",
	Short: "comer",
	Long:  `comer --module=[module] --path=[dir]`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {

		comerIns := comer.NewComer()
		comerIns.Start(cmd, args)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}

}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.comer.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().String(`module`, ``, `go.mod module(go.mod文件的module名称)`)
	rootCmd.PersistentFlags().String(`path`, ``, `project root (项目所在目录)`)
}
