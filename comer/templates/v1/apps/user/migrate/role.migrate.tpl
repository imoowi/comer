package migrate

import (
	"{{.moduleName}}/apps/user/models"
	"{{.moduleName}}/components"
	"{{.moduleName}}/global"
)

type RoleMigrate struct {
	db *components.MysqlODM
}

func newRoleMigrate() *RoleMigrate {
	return &RoleMigrate{
		db: global.MysqlDb,
	}
}
func init() {
	global.RegisterMigrateContainerProviders(doRoleMigrate)
}
func doRoleMigrate() {
	r := newRoleMigrate()
	r.db.Client.Set("gorm:table_options", "ENGINE=InnoDB,COMMENT='角色表'").AutoMigrate(&models.Role{})
}