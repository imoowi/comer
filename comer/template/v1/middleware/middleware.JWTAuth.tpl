/*
Copyright © 2023 yuanjun<simpleyuan@gmail.com>
*/
package middlewares

import (
	"net/http"
	"strings"
	"time"

	"{{.moduleName}}/apps/user/services"
	token "{{.moduleName}}/middlewares/token"
	"{{.moduleName}}/utils/response"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {

		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// 这里的具体实现方式要依据你的实际业务情况决定
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			response.Error("请求头中auth为空", http.StatusUnauthorized, c)
			c.Abort()
			return
		}
		//判断用户是否主动注销过
		isLogouted := services.User.IsLogouted(c, authHeader)
		if isLogouted {
			response.Error("token 失效", http.StatusUnauthorized, c)
			c.Abort()
			return
		}

		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			response.Error("请求头中auth为空", http.StatusUnauthorized, c)
			c.Abort()
			return
		}

		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := token.ParseToken(parts[1])
		if err != nil {
			response.Error("invaled Token", http.StatusUnauthorized, c)
			c.Abort()
			return
		}
		//根据admin_id查询用户是否存在
		user, err := services.User.One(c, mc.UserId)
		if err != nil {
			response.Error(err.Error(), http.StatusUnauthorized, c)
			c.Abort()
			return
		}
		if user.ID <= 0 {
			response.Error(`user dose not exist`, http.StatusUnauthorized, c)
			c.Abort()
			return
		}
		// 判断token是否即将过期
		refreshTokenTimeout := viper.GetDuration("settings.jwt.refresh_token_timeout")
		refreshTokenTimeoutSeconds := refreshTokenTimeout.Seconds()
		invalidTimeout := mc.StandardClaims.ExpiresAt - int64(refreshTokenTimeoutSeconds)
		timeNow := time.Now().Unix()
		if invalidTimeout < int64(timeNow) {
			//返回新token
			tokenString, _ := token.GenToken(mc.Username, mc.UserId)
			c.Header(`token`, tokenString)
		}
		//
		// 将当前请求的username信息保存到请求的上下文c上
		c.Set("username", mc.Username) // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
		c.Set("uid", mc.UserId)

		isSuperAdmin := services.User.IsSuper(c, mc.UserId)
		c.Set("isSuperAdmin", isSuperAdmin)
		c.Set(`isInit`, false)

		c.Next() 
	}
}
