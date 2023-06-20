/*
Copyright © 2023 yuanjun<simpleyuan@gmail.com>
*/
package global

import (
	"log"

	"{{.moduleName}}/components"
)

var MysqlDb *components.MysqlODM

// 初始化mysql
func initMysql() {
	// 获取logger相关的配置信息
	config := Config.Sub("mysql")
	var mysqlConfig *components.MysqlODMConfig
	err := config.Unmarshal(&mysqlConfig)
	if err != nil {
		log.Fatal(err)
	}
	mode := Config.GetString("application.mode")
	mysqlConfig.Mode = mode
	MysqlDb = components.NewMysqlODM(mysqlConfig)
}
