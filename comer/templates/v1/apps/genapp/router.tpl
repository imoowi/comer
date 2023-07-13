/*
Copyright © 2023 yuanjun<simpleyuan@gmail.com>
*/
package {{.appName}}

import (
	"github.com/gin-gonic/gin"
	_ "{{.ModuleName}}/apps/{{.appName}}/migrates"
	"{{.ModuleName}}/middlewares"
	"{{.ModuleName}}/apps/{{.appName}}/handlers"
	"{{.ModuleName}}/router"
)

func init() {
	router.RegisterRoute(Routers)
}

func Routers(e *gin.Engine) {
	api := e.Group("/api")
	api.Use(middlewares.JWTAuthMiddleware())
	api.Use(middlewares.CasbinMiddleware())
	{{.lHandlerName}}s := api.Group("/{{.handlerName2Dash}}s")
	{
		{{.lHandlerName}}s.GET("", handlers.{{.HandlerName}}PageList) //分页
		{{.lHandlerName}}s.GET("/:id", handlers.{{.HandlerName}}One) //一个
		{{.lHandlerName}}s.POST("", handlers.{{.HandlerName}}Add) //新增
		{{.lHandlerName}}s.PUT("/:id", handlers.{{.HandlerName}}Update) //更新
		{{.lHandlerName}}s.PATCH("/:id", handlers.{{.HandlerName}}Patch) //部分更新
		{{.lHandlerName}}s.DELETE("/:id", handlers.{{.HandlerName}}Del) //默认为软删除
	}

	//!import:do-not-delete-this-line,不要删除此行，主要用于代码生成器
}
