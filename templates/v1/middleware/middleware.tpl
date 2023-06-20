package middlewares

import (
	"github.com/gin-gonic/gin"
)

func InitMiddleware(r *gin.Engine) {

	r.Use(LoggerMiddleware())
	r.Use(CorsMiddleware())
	r.Use(RateLimitMiddleware()) //初始100个令牌，每秒允许100个令牌通过
}
