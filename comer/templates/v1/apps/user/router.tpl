package user

import (
	"github.com/gin-gonic/gin"
	"{{.moduleName}}/apps/user/handlers"
	_ "{{.moduleName}}/apps/user/migrates"
	"{{.moduleName}}/middlewares"
	"{{.moduleName}}/router"
)

func init() {
	router.RegisterRoute(Routers)
}

func Routers(e *gin.Engine) {
	auth := e.Group(`/api`)
	{
		auth.POST(`/login`, middlewares.VcodeMiddleware(), handlers.AuthLogin)
		auth.GET(`/logout`, middlewares.JWTAuthMiddleware(), handlers.AuthLogout)
		auth.POST(`/chpwd`, middlewares.VcodeMiddleware(), middlewares.JWTAuthMiddleware(), handlers.AuthChgPwd)
	}
}
