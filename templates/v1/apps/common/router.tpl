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
