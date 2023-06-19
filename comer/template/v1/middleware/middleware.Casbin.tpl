package middleware

import (
	"fmt"
	"github.com/imoowi/commer/utils/response"
	"net/http"
	"strings"
	"{{.moduleName}}/global"
	"{{.moduleName}}/services"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func CasbinMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		username := c.GetString("username")
		if username == `` {
			// fmt.Println(`c.Get("username") is: `, username)
			response.Error(`need login ...`, http.StatusUnauthorized, c)
			c.Abort()
			return
		}

		// roleId := c.GetString("role_id")
		adminId := c.GetString(`admin_id`)
		roleId := services.AuthRoleId(adminId)
		if roleId == `` {
			response.Error(`need login ...`, http.StatusUnauthorized, c)
			c.Abort()
			return
		}

		e := global.Casbin
		// fmt.Println(`admin.roleId=`, roleId, `v1=`, c.Request.URL.Path, `v2=`, c.Request.Method)
		canAccess := false
		// _roleId := cast.ToString(roleId)
		roleSlice := strings.Split(roleId, ",")
		for _, _roleId := range roleSlice {
			// Check the permission.
			ok, err := e.Enforce(_roleId, c.Request.URL.Path, c.Request.Method)
			if err != nil {
				fmt.Printf(`roleid:%v, has no permisson:%v\n`, _roleId, err.Error())
				continue
			}
			if !ok {
				fmt.Println(`check role perm  is error:`, ok)
				continue
			}
			canAccess = true
			break
		}
		if !canAccess {
			response.Error(`您没有访问的权限`, http.StatusForbidden, c)
			c.Abort()
			return
		}
		c.Next()
		/*
			// Check the permission.
			ok2, err2 := e.Enforce(username, c.Request.URL.Path, c.Request.Method)
			if err2 != nil {
				fmt.Println(`casbin:`, err.Error())
				response.Error(`您没有访问的权限3`, http.StatusBadRequest, c)
				c.Abort()
				return
			}

			if !ok2 {
				fmt.Println(`ok is:`, ok2)
				response.Error(`您没有访问的权限4`, http.StatusBadRequest, c)
				c.Abort()
				return
			}

			//*/

	}
}
