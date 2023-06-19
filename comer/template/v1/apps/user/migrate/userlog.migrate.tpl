package migrate

import (
	"{{.moduleName}}/apps/user/models"
	"{{.moduleName}}/global"
)

type UserLogMigrate struct {
	db *global.MysqlODM
}

func newUserLogMigrate() *UserLogMigrate {
	return &UserLogMigrate{
		db: global.MysqlDb,
	}
}
func init() {
	global.RegisterMigrateContainerProviders(doUserLogMigrate)
}
func doUserLogMigrate() {
	r := newUserLogMigrate()
	r.db.Client.Set("gorm:table_options", "ENGINE=InnoDB,COMMENT='用户行为记录表'").AutoMigrate(&models.UserLog{})
}
