package global

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
)

// 全局权限
var Casbin *casbin.Enforcer

// 初始化casbin
func initCasbin() error {
	dataSourceName := MysqlDb.Config.Casbin
	a, _ := gormadapter.NewAdapter("mysql", dataSourceName, true) // Your driver and data source.
	e, err := casbin.NewEnforcer("config/casbin.conf", a)
	if err != nil {
		return err
	}
	// Load the policy from DB.
	err = e.LoadPolicy()
	if err != nil {
		return err
	}
	Casbin = e
	return nil
}
