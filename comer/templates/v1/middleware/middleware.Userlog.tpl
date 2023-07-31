package middlewares

import (
	"{{.moduleName}}/apps/user/models"
	"{{.moduleName}}/apps/user/services"
	"{{.moduleName}}/global"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

// 用户日志中间件
func UserLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// user_log_obj_id := c.GetUint(`user_log_obj_id`) //操作对象id
		// user_log_res_type, ok := c.Get(`user_log_res_type`) //操作对象的资源类型
		user_log_obj_id := 0
		if data, ok := c.Get(`user_log_obj_id`); ok {
			user_log_obj_id = cast.ToInt(data)
		} else {
			c.Abort()
			return
		}
		user_log_res_type_string := ``
		if user_log_res_type, ok := c.Get(`user_log_res_type`); ok {
			uid := c.GetUint(`uid`)
			if uid > 0 {
				if data, ok2 := global.ResTypes[user_log_res_type.(global.RES_TYPE)]; ok2 {
					user_log_res_type_string = data
				}
				action := `-`
				switch c.Request.Method {
				case `GET`:
					action = `访问`
				case `POST`:
					action = `新建`
				case `UPDATE`:
					fallthrough
				case `PATCH`:
					fallthrough
				case `PUT`:
					action = `修改`
				case `DELETE`:
					action = `删除`
				}
				userlog := &models.UserLog{
					UserID:     uid,
					LogType:    global.USER_LOG_TYPE_ADD,
					ResType:    user_log_res_type.(global.RES_TYPE),
					LogContent: `用户【` + c.GetString(`username`) + `】` + action + `了` + user_log_res_type_string + `【` + cast.ToString(user_log_obj_id) + `】`,
					IP:         c.ClientIP(),
				}
				services.UserLog.Add(c, userlog)
			}
			c.Next()
		}
	}
}
