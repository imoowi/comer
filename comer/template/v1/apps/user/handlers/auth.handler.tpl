package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"{{.moduleName}}/apps/user/models"
	"{{.moduleName}}/apps/user/services"
	"{{.moduleName}}/middlewares/token"
	"{{.moduleName}}/utils/response"
)

// @Summary	登录
// @Tags		认证
// @Accept		application/json
// @Produce	application/json
// @Param		body	body	models.UserLogin	true	"登录信息"
// @Success	200
// @Failure	400	{object}	string	"请求错误"
// @Failure	401	{object}	string	"token验证失败"
// @Failure	500	{object}	string	"内部错误"
// @Router		/api/auth/login [post]
func AuthLogin(c *gin.Context) {
	var userLogin *models.UserLogin
	err := c.ShouldBindJSON(&userLogin)
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest, c)
		return
	}
	user, err := services.User.Login(c, userLogin)
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest, c)
		return
	}
	if user == nil || user.ID <= 0 {
		response.Error(err.Error(), http.StatusBadRequest, c)
		return
	}
	// 生成Token
	tokenString, _ := token.GenToken(user.Username, user.ID)

	user.Passwd = ``
	user.Salt = ``
	response.OK(gin.H{"token": tokenString, "info": user}, c)
}

// @Summary	退出
// @Tags		认证
// @Accept		application/json
// @Produce	application/json
// @Param		Authorization	header	string	true	"Bearer 用户令牌"
// @Success	200
// @Failure	400	{object}	string	"请求错误"
// @Failure	401	{object}	string	"token验证失败"
// @Failure	500	{object}	string	"内部错误"
// @Router		/api/auth/logout [get]
func AuthLogout(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		response.Error("请求头中auth为空", http.StatusUnauthorized, c)
		return
	}
	ok := services.User.Logout(c, authHeader)
	response.OK(ok, c)
}

// @Summary	改密
// @Tags		认证
// @Accept		application/json
// @Produce	application/json
// @Param		Authorization	header	string				true	"Bearer 用户令牌"
// @Param		body			body	models.UserChgPwd	true	"改密信息"
// @Success	200
// @Failure	400	{object}	string	"请求错误"
// @Failure	401	{object}	string	"token验证失败"
// @Failure	500	{object}	string	"内部错误"
// @Router		/api/auth/chgpwd [post]
func AuthChgPwd(c *gin.Context) {
	var userChgPwd *models.UserChgPwd
	err := c.ShouldBindJSON(&userChgPwd)
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest, c)
		return
	}
	newJwtToken, err := services.User.ChgPwd(c, userChgPwd)
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest, c)
		return
	}
	if newJwtToken != `` {
		c.Header(`token`, newJwtToken)
	}
	response.OK(newJwtToken, c)
}
