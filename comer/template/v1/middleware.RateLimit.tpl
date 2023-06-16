package middleware

import (
	"github.com/imoowi/commer/util/response"
	"{{.moduleName}}/global"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

func RateLimitMiddleware() gin.HandlerFunc {
	bucket := ratelimit.NewBucketWithQuantum(time.Second, global.Viper.GetInt64(`settings.ratelimit.cap`), global.Viper.GetInt64(`settings.ratelimit.quantum`))
	return func(c *gin.Context) {
		if bucket.TakeAvailable(1) < 1 {
			response.Error(`rate limit ...`, http.StatusForbidden, c)
			c.Abort()
			return
		}
		c.Next()
	}
}
