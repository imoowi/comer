/*
Copyright © 2023 yuanjun<simpleyuan@gmail.com>
*/
package global

import (
	"errors"
	"log"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Config *viper.Viper
var (
	ErrConfigIsNotAbsPath = errors.New("配置文件不是绝对路径")
)

func InitConfig(configPath string) {
	cfg, err := getConfig(configPath)
	if err != nil {
		panic(err)
	}
	cfg.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := cfg.ReadInConfig(); err == nil {
		log.Println("Using config file:", cfg.ConfigFileUsed())
	}
	cfg.OnConfigChange(func(e fsnotify.Event) {
		log.Println(`config file changed:`, e.Name)
		err := cfg.ReadInConfig()
		if err != nil {
			panic(err)
		}
	})
	go cfg.WatchConfig()
	Config = cfg
}

// @configPath	string 配置文件路径，绝对路径
func getConfig(configPath string) (*viper.Viper, error) {
	if !filepath.IsAbs(configPath) {
		return nil, ErrConfigIsNotAbsPath
	}
	// dir := filepath.Dir(configPath)
	// fileName := filepath.Base(configPath)
	config := viper.New()
	config.SetConfigType("yml")
	config.SetConfigFile(configPath)
	return config, nil
}
