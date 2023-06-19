package global

func Bootstrap() {
	initLog()
	initMysql()
	initRedis()
	initCasbin()
}
