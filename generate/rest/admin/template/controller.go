/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-04-07 16:12:19
 */

package template

import (
	"framework/app"
	"framework/proto/request"
	"framework/proto/response"
	"framework/proto/result"
	"generate/service/template"
	"github.com/gin-gonic/gin"
)

// @Summary  获取某一个
// @Tags     ${Template.name}
// @Accept   json
// @Produce  json
// @Success  200  {object}  result.Result{data=Template}  true
// @Router   /admin/Template/:id [get]
func getOne(ctx *gin.Context) {
	// 获取 id
	id := ctx.Param("id")

	// 根据 id 查询
	one := template.FindById(id)

	// 未找到结果
	if one == nil {
		response.New(ctx).Data(result.NotFound)
		return
	}

	// 返回结果
	response.New(ctx).Data(one)
}

// @Summary  获取所有
// @Tags     ${Template.name}
// @Accept   json
// @Produce  json
// @Param    page      query     int                                 true  "分页，从 0 开始"
// @Param    pageSize  query     int                                 true  "分页大小"
// @Success  200       {object}  result.PageResult{data=[]Template}  true
// @Router   /admin/Template [get]
func getAll(ctx *gin.Context) {
	// 分页请求对象
	var req request.PageRequest

	// 较验参数
	app.ValidateParameter(ctx, &req)

	// 根据 userId 查询所有
	array, pagination := template.FindAll(req.Page, req.PageSize)

	// 未找到结果
	if len(array) == 0 {
		response.New(ctx).Data(result.NotFound)
		return
	}

	// 返回结果
	response.New(ctx).Data(result.PageResult{
		Result:     result.Ok.WithData(array),
		Pagination: pagination,
	})
}

// @Summary  更新
// @Tags     ${Template.name}
// @Accept   json
// @Produce  json
// @Param    json  body      putTemplateReq  true  "struct"
// @Success  200   {object}  result.Result   true
// @Router   /admin/Template/:id [put]
func putOne(ctx *gin.Context) {
	// 分页请求对象
	var req putTemplateReq

	// 较验参数
	app.ValidateParameter(ctx, &req)

	// 更新
	one := template.Template{
		Name: req.Name,
	}
	template.UpdateById(one.Id, one)

	// 返回结果
	response.New(ctx).Ok()
}

// @Summary  删除
// @Tags     ${Template.name}
// @Accept   json
// @Produce  json
// @Success  200  {object}  result.Result  true
// @Router   /admin/Template/ [put]
func deleteOne(ctx *gin.Context) {
	// 获取 id
	id := ctx.Param("id")

	// 删除
	template.DeleteById(id)

	// 返回结果
	response.New(ctx).Ok()
}
