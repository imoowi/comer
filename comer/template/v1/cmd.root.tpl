/*
Copyright © 2023 yuanjun<simpleyuan@gmail.com>
*/
package cmd

import (
	"os"
	"path/filepath"
	"{{.moduleName}}/global"

	"github.com/spf13/cobra"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "{{.moduleName}}",
	Short: "{{.moduleName}}",
	Long:  ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
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
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is configs/settings-local.yml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// 未指定配置文件时，构造默认的配置文件路径
	if cfgFile == "" || !filepath.IsAbs(cfgFile) {
		work, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		if cfgFile == "" {
			cfgFile = "configs/settings-local.yml"
		}
		cfgFile = filepath.Join(work, cfgFile)
	}
	_, err := os.Stat(cfgFile)
	if err != nil {
		panic(err)
	}
	// 初始化配置文件
	global.InitConfig(cfgFile)
}
