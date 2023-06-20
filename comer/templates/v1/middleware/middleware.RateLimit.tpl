/*
Copyright © 2023 yuanjun<simpleyuan@gmail.com>
*/
package middlewares

import (
	"net/http"
	"time"

	"{{.moduleName}}/global"
	"{{.moduleName}}/utils/response"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

func RateLimitMiddleware() gin.HandlerFunc {
	bucket := ratelimit.NewBucketWithQuantum(time.Second, global.Config.GetInt64(`ratelimit.cap`), global.Config.GetInt64(`ratelimit.quantum`))
	return func(c *gin.Context) {
		if bucket.TakeAvailable(1) < 1 {
			response.Error(`rate limit ...`, http.StatusForbidden, c)
			c.Abort()
			return
		}
		c.Next()
	}
}
