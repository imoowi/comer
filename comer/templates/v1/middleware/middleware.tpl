/*
generated by comer,https://github.com/imoowi/comer
Copyright © 2023 jun<simpleyuan@gmail.com>
*/
package middlewares

import (
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

func InitMiddleware(r *gin.Engine) {

	r.Use(LoggerMiddleware())
	r.Use(CorsMiddleware())
	r.Use(RateLimitMiddleware()) //初始100个令牌，每秒允许100个令牌通过
	r.Use(requestid.New())
	r.Use(RequestIdMiddleware())
}
