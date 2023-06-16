package global

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlODMConfig struct {
	Dsn    string `json:"dsn"`    //链接地址
	Casbin string `json:"casbin"` //casbin
	Mode   string `json:"mode"`
}

type MysqlODM struct {
	Config *MysqlODMConfig
	Client *gorm.DB
}

// 构造函数
func NewMysqlODM(config *MysqlODMConfig) *MysqlODM {
	var sqlDb *sql.DB
	dsn := config.Dsn
	_db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	sqlDb, err = _db.DB()
	if err != nil {
		fmt.Println(`sqlDB=`, sqlDb)
		panic("failed to connect mysql database")
	} else {
		fmt.Println(`Connected to MySql!`)
		sqlDb.SetMaxIdleConns(10)
		sqlDb.SetMaxOpenConns(100)
		sqlDb.SetConnMaxLifetime(time.Hour)
	}
	var db *gorm.DB

	if config.Mode == `dev` {
		db = _db.Debug()
	} else {
		db = _db
	}

	odm := MysqlODM{
		Config: config,
		Client: db,
	}
	return &odm
}

var MysqlDb *MysqlODM

// 初始化mysql
func initMysql() {
	// 获取logger相关的配置信息
	config := Config.Sub("mysql")
	var mysqlConfig *MysqlODMConfig
	err := config.Unmarshal(&mysqlConfig)
	if err != nil {
		log.Fatal(err)
	}
	mode := Config.GetString("application.mode")
	mysqlConfig.Mode = mode
	MysqlDb = NewMysqlODM(mysqlConfig)
}