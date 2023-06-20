package user

import (
	"github.com/gin-gonic/gin"
	"{{.moduleName}}/apps/user/handlers"
	_ "{{.moduleName}}/apps/user/migrate"
	"{{.moduleName}}/middlewares"
	"{{.moduleName}}/router"
)

func init() {
	router.RegisterRoute(Routers)
}

func Routers(e *gin.Engine) {
	auth := e.Group(`/api/auth`)
	{
		auth.POST(`/login`, middlewares.VcodeMiddleware(), handlers.AuthLogin)
		auth.GET(`/logout`, middlewares.JWTAuthMiddleware(), handlers.AuthLogout)
		auth.POST(`/chgpwd`, middlewares.JWTAuthMiddleware(), middlewares.VcodeMiddleware(), handlers.AuthChgPwd)
	}
}
