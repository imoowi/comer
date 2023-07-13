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
		auth.POST(`/auth-login`, middlewares.VcodeMiddleware(), handlers.AuthLogin)
		auth.GET(`/auth-logout`, middlewares.JWTAuthMiddleware(), handlers.AuthLogout)
		auth.POST(`/auth-chpwd`, middlewares.VcodeMiddleware(), middlewares.JWTAuthMiddleware(), handlers.AuthChgPwd)
	}
}
