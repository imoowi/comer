package migrate

import (
	"{{.moduleName}}/apps/user/models"
	"{{.moduleName}}/components"
	"{{.moduleName}}/global"
)

type UserMigrate struct {
	db *components.MysqlODM
}

func newUserMigrate() *UserMigrate {
	return &UserMigrate{
		db: global.MysqlDb,
	}
}
func init() {
	global.RegisterMigrateContainerProviders(doUserMigrate)
}
func doUserMigrate() {
	r := newUserMigrate()
	r.db.Client.Set("gorm:table_options", "ENGINE=InnoDB,COMMENT='用户表'").AutoMigrate(&models.User{})
}
