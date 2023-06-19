package router

import (
	"{{.moduleName}}/common/utils/response"
	"{{.moduleName}}/middlewares"

	"github.com/gin-gonic/gin"
)

type Router func(*gin.Engine)

var router = []Router{}

func InitRouter() *gin.Engine {
	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	middleware.InitMiddleware(r)

	for _, route := range router {
		route(r)

	}
	r.GET(`/api/casbins/allapi`, func(ctx *gin.Context) {

		// @Router       /api/casbins/allapi [get]
		routers := r.Routes()
		allapi := make([]map[string]string, 0)
		for _, v := range routers {
			allapi = append(allapi, map[string]string{"method": v.Method, "path": v.Path})
		}
		response.OK(allapi, ctx)
	})
	return r
}

// router  其余模块在init中调用RegisterRoute注册
func RegisterRoute(r ...Router) {
	router = append(router, r...)
}
