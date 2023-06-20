package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"{{.moduleName}}/components"
	"{{.moduleName}}/utils/response"
)

// @Summary	验证码  /api/common/captcha?type=audio|string|math|chinese
// @Tags		公用
// @Accept		application/json
// @Produce	application/json
// @Param		type	query	string	false	"验证码类型：audio|string|math|chinese；默认为 math"
// @Success	200
// @Failure	400	{object}	string	"请求错误"
// @Failure	401	{object}	string	"token验证失败"
// @Failure	500	{object}	string	"内部错误"
// @Router		/api/common/captcha [get]
func Captcha(c *gin.Context) {
	driverType := c.DefaultQuery(`type`, `digit`)
	id, b64s, err := components.GenCaptcha(driverType)
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest, c)
		return
	}
	captchaMap := map[string]any{
		`captcha_id`:   id,
		`captcha_code`: b64s,
	}
	response.OK(captchaMap, c)
}
