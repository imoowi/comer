/*
generated by comer,https://github.com/imoowi/comer
Copyright © 2023 jun<simpleyuan@gmail.com>
*/
package common

import (
	"github.com/gin-gonic/gin"
	"{{.moduleName}}/apps/common/handlers"
	"{{.moduleName}}/router"
)

func init() {
	router.RegisterRoute(Routers)
}

func Routers(e *gin.Engine) {
	auth := e.Group(`/api/common`)
	{
		auth.GET(`/captcha`, handlers.Captcha)
	}
}
