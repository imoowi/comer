package middleware

import (
	"github.com/imoowi/commer/utils/response"
	"{{.moduleName}}/global"
	token "{{.moduleName}}/middlewares/token"
	"{{.moduleName}}/services"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		//添加微信过来的接口访问权限判断
		wxAuthHeader := c.Request.Header.Get("Authorization-wx")
		if wxAuthHeader != `` {
			//判断用户是否主动注销过
			isWxLogouted := service.IsUserLogouted(wxAuthHeader)
			if isWxLogouted {
				response.Error("token 失效", http.StatusUnauthorized, c)
				c.Abort()
				return
			}
			// 按空格分割
			wxParts := strings.SplitN(wxAuthHeader, " ", 2)
			if len(wxParts) == 2 && wxParts[0] == "Bearer" {
				// wxParts[1]是获取到的wxTokenString，我们使用之前定义好的解析JWT的函数来解析它
				wxMc, err := token.ParseWxUserToken(wxParts[1])
				if err == nil {
					// 判断token是否即将过期
					wxRefreshTokenTimeout := viper.GetDuration("settings.jwt.refresh_token_timeout")
					wxRefreshTokenTimeoutSeconds := wxRefreshTokenTimeout.Seconds()
					wxInvalidTimeout := wxMc.StandardClaims.ExpiresAt - int64(wxRefreshTokenTimeoutSeconds)
					wxTimeNow := time.Now().Unix()
					if wxInvalidTimeout < int64(wxTimeNow) {
						//返回新token
						wxTokenString, _ := token.GenWxUserToken(wxMc.WxUserId)
						c.Header(`wx-token`, wxTokenString)
					}
					// 将当前请求的username信息保存到请求的上下文c上
					c.Set("wx_user_id", wxMc.WxUserId)

					if wxMc.WxUserId != "" {
						wxUser := service.WechatMiniProgramUserInfo(wxMc.WxUserId)
						adminUser, _ := service.GetAdmin(wxUser.SystemUserId)
						if adminUser.ID != "" {
							c.Set("username", adminUser.Username)
							c.Set("admin_id", adminUser.ID)
							c.Set("role_id", service.AuthRoleId(adminUser.ID))
						}
					}

					c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
					return
				}

			}
		}

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
		isLogouted := service.IsUserLogouted(authHeader)
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
		_, err = service.GetAdmin(mc.AdminId)
		if err != nil {
			response.Error(err.Error(), http.StatusUnauthorized, c)
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
			tokenString, _ := token.GenToken(mc.Username, mc.AdminId, mc.RoleId)
			c.Header(`token`, tokenString)
		}
		//
		// 将当前请求的username信息保存到请求的上下文c上
		c.Set("username", mc.Username)
		c.Set("admin_id", mc.AdminId)
		c.Set("role_id", mc.RoleId)

		isSuperAdmin := false
		if mc.RoleId != `` {
			roles := strings.Split(mc.RoleId, `,`)
			for _, v := range roles {
				roleMap := service.GetOneRole(v)
				// if roleMap.Name == `超级管理员` {
				if roleMap.Name == global.Viper.GetString(`system.superAdminRole`) {
					isSuperAdmin = true
				}
			}
		}
		c.Set("is_super_admin", isSuperAdmin)

		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}
