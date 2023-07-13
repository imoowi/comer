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
	api := e.Group(`/api`)
	{
		api.POST(`/auth-login`, middlewares.VcodeMiddleware(), handlers.AuthLogin)
		api.GET(`/auth-logout`, middlewares.JWTAuthMiddleware(), handlers.AuthLogout)
		api.POST(`/auth-chpwd`, middlewares.VcodeMiddleware(), middlewares.JWTAuthMiddleware(), handlers.AuthChgPwd)
	}
	
	//!import:do-not-delete-this-line,不要删除此行，主要用于代码生成器
}
