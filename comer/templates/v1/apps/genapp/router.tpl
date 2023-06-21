/*
Copyright © 2023 yuanjun<simpleyuan@gmail.com>
*/
package {{.appName}}

import (
	"github.com/gin-gonic/gin"
	_ "{{.ModuleName}}/apps/{{.handlerName}}/migrates"
	"{{.ModuleName}}/middlewares"
	"{{.ModuleName}}/apps/{{.appName}}/handlers"
	"{{.ModuleName}}/router"
)

func init() {
	router.RegisterRoute(Routers)
}

func Routers(e *gin.Engine) {
	{{.appName}} := e.Group("/api/{{.appName}}")
	{{.appName}}.Use(middlewares.JWTAuthMiddleware())
	{{.appName}}.Use(middlewares.CasbinMiddleware())
	{{.handlerName}}s := {{.appName}}.Group("/{{.handlerName}}s")
	{
		{{.handlerName}}s.GET("", handlers.{{.HandlerName}}PageList) //分页
		{{.handlerName}}s.GET("/:id", handlers.{{.HandlerName}}One) //一个
		{{.handlerName}}s.POST("", handlers.{{.HandlerName}}Add) //新增
		{{.handlerName}}s.PUT("/:id", handlers.{{.HandlerName}}Update) //更新
		{{.handlerName}}s.DELETE("/:id", handlers.{{.HandlerName}}Del) //默认为软删除
	}

	//!import:do-not-delete-this-line,不要删除此行，主要用于代码生成器
}
