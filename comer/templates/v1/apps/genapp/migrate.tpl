/*
Copyright © 2023 yuanjun<simpleyuan@gmail.com>
*/
package migrate

import (
	"{{.ModuleName}}/apps/{{.appName}}/models"
	"{{.ModuleName}}/components"
	"{{.ModuleName}}/global"
)

type {{.ModelName}}Migrate struct {
	db *components.MysqlODM
}

func new{{.ModelName}}Migrate() *{{.ModelName}}Migrate {
	return &{{.ModelName}}Migrate{
		db: global.MysqlDb,
	}
}
func init() {
	global.RegisterMigrateContainerProviders(do{{.ModelName}}Migrate)
}
func do{{.ModelName}}Migrate() {
	r := new{{.ModelName}}Migrate()
	r.db.Client.Set("gorm:table_options", "ENGINE=InnoDB,COMMENT='{{.ModelName}}表'").AutoMigrate(&models.{{.ModelName}}{})
}