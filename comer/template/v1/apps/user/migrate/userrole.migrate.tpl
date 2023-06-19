package migrate

import (
	"{{.moduleName}}/apps/user/models"
	"{{.moduleName}}/global"
)

type UserRoleMigrate struct {
	db *global.MysqlODM
}

func newUserRoleMigrate() *UserRoleMigrate {
	return &UserRoleMigrate{
		db: global.MysqlDb,
	}
}
func init() {
	global.RegisterMigrateContainerProviders(doUserRoleMigrate)
}
func doUserRoleMigrate() {
	r := newUserRoleMigrate()
	r.db.Client.Set("gorm:table_options", "ENGINE=InnoDB,COMMENT='用户角色关系表'").AutoMigrate(&models.UserRole{})
}
