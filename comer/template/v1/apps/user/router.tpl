package user

import (
	"github.com/gin-gonic/gin"
	"{{.moduleName}}/apps/user/handlers"
	"{{.moduleName}}/router"
)

func init() {
	router.RegisterRoute(Routers)
}

func Routers(e *gin.Engine) {
	auth := e.Group(`/api/auth`)
	{
		auth.POST(`/login`, handlers.AuthLogin)
		auth.GET(`/logout`, handlers.AuthLogout)
		auth.POST(`/chgpwd`, handlers.AuthChgPwd)
	}
}
