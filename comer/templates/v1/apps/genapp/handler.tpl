/*
Copyright © 2023 yuanjun<simpleyuan@gmail.com>
*/
package handlers

import (
	"github.com/gin-gonic/gin"
	"{{.ModuleName}}/utils/response"
	"{{.ModuleName}}/apps/{{.appName}}/models"
	"{{.ModuleName}}/apps/{{.appName}}/services"
	"{{.ModuleName}}/utils/request"
	"net/http"

	"github.com/spf13/cast"
)

// @Summary	分页列表(pagelist)
// @Tags		{{.SwaggerTags}}
// @Accept		application/json
// @Produce	application/json
// @Param		Authorization	header		string				true	"Bearer 用户令牌"
// @Param 		page 			query		number				false	"页码,默认为1"
// @Param 		pageSize 		query		number				false	"页数,默认为20，最小为1，最大不超过1000"
// @Success	200				{object}	response.PageList	"成功"
// @Failure	400				{object}	string				"请求错误"
// @Failure	401				{object}	string				"token验证失败"
// @Failure	500				{object}	string				"内部错误"
// @Router		/api/{{.handlerName2Dash}}s [get]
func {{.HandlerName}}PageList(c *gin.Context) {
	var req request.PageList
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest,c)
		return
	}
	if 0 >= req.Page{
		req.Page = 1
	}
	if 0 >= req.PageSize{
		req.PageSize = 20
	}
	if req.PageSize >= 1000 {
		req.PageSize = 1000
	}
	result, err := services.{{.ServiceName}}.PageList(c, &req)
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest,c)
		return
	}
	response.OK(result,c)
}

// @Summary	详情(one)
// @Tags		{{.SwaggerTags}}
// @Accept		application/json
// @Produce	application/json
// @Param		Authorization	header	string	true	"Bearer 用户令牌"
// @Param		id				path	int		true	"id"
// @Success	200
// @Failure	400	{object}	string	"请求错误"
// @Failure	401	{object}	string	"token验证失败"
// @Failure	500	{object}	string	"内部错误"
// @Router		/api/{{.handlerName2Dash}}s/:id [get]
func {{.HandlerName}}One(c *gin.Context) {
	id := c.Param(`id`)
	if id == `` {
		response.Error(`pls input id`, http.StatusBadRequest,c)
		return
	}
	one, err := services.{{.ServiceName}}.One(c, cast.ToUint(id))
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest,c)
		return
	}
	response.OK(one,c)
}

// @Summary	新增(add)
// @Tags		{{.SwaggerTags}}
// @Accept		application/json
// @Produce	application/json
// @Param		Authorization	header	string	true	"Bearer 用户令牌"
// @Param		id				path	int		true	"id"
// @Success	200
// @Failure	400	{object}	string	"请求错误"
// @Failure	401	{object}	string	"token验证失败"
// @Failure	500	{object}	string	"内部错误"
// @Router		/api/{{.handlerName2Dash}}s [post]
func {{.HandlerName}}Add(c *gin.Context) {
	var {{.modelName}} *models.{{.ModelName}}
	err := c.ShouldBindJSON(&{{.modelName}})
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest,c)
		return
	}
	newId, err := services.{{.ServiceName}}.Add(c, {{.modelName}})
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest,c)
		return
	}
	response.OK(newId,c)
}

// @Summary	更新(update)
// @Tags		{{.SwaggerTags}}
// @Accept		application/json
// @Produce	application/json
// @Param		Authorization	header	string	true	"Bearer 用户令牌"
// @Param		id				path	int		true	"id"
// @Success	200
// @Failure	400	{object}	string	"请求错误"
// @Failure	401	{object}	string	"token验证失败"
// @Failure	500	{object}	string	"内部错误"
// @Router		/api/{{.handlerName2Dash}}s/:id [put]
func {{.HandlerName}}Update(c *gin.Context) {
	id := c.Param(`id`)
	if id == `` {
		response.Error(`pls input id`, http.StatusBadRequest,c)
		return
	}
	var {{.modelName}} *models.{{.ModelName}}
	err := c.ShouldBindJSON(&{{.modelName}})
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest,c)
		return
	}
	updated, err := services.{{.ServiceName}}.Update(c, {{.modelName}}, cast.ToUint(id))
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest,c)
		return
	}
	response.OK(updated,c)
}

// @Summary	部分更新(update)
// @Tags		{{.SwaggerTags}}
// @Accept		application/json
// @Produce	application/json
// @Param		Authorization	header	string	true	"Bearer 用户令牌"
// @Param		id				path	int		true	"id"
// @Success	200
// @Failure	400	{object}	string	"请求错误"
// @Failure	401	{object}	string	"token验证失败"
// @Failure	500	{object}	string	"内部错误"
// @Router		/api/{{.handlerName2Dash}}s/:id [patch]
func {{.HandlerName}}Patch(c *gin.Context) {
	id := c.Param(`id`)
	if id == `` {
		response.Error(`pls input id`, http.StatusBadRequest, c)
		return
	}
	patchObj := make(map[string]any)
	err := c.BindJSON(&patchObj)
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest, c)
		return
	}
	updated, err := services.{{.ServiceName}}.PatchUpdate(c, patchObj, cast.ToUint(id))
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest, c)
		return
	}
	response.OK(updated, c)
}

// @Summary	删除(delete)
// @Tags		{{.SwaggerTags}}
// @Accept		application/json
// @Produce	application/json
// @Param		Authorization	header	string	true	"Bearer 用户令牌"
// @Param		id				path	int		true	"id"
// @Success	200
// @Failure	400	{object}	string	"请求错误"
// @Failure	401	{object}	string	"token验证失败"
// @Failure	500	{object}	string	"内部错误"
// @Router		/api/{{.handlerName2Dash}}s/:id [delete]
func {{.HandlerName}}Del(c *gin.Context) {
	id := c.Param(`id`)
	if id == `` {
		response.Error(`pls input id`, http.StatusBadRequest,c)
		return
	}
	deleted, err := services.{{.ServiceName}}.Delete(c, cast.ToUint(id))
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest,c)
		return
	}
	response.OK(deleted,c)
}
