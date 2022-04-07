/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-04-07 16:12:11
 */

package template

import (
	"github.com/gin-gonic/gin"
	"medium-server-go/framework/app"
	"medium-server-go/framework/low_code/service/template"
	"medium-server-go/framework/proto/request"
	"medium-server-go/framework/proto/response"
	"medium-server-go/framework/proto/result"
	"medium-server-go/framework/proto/token"
)

// @Summary  获取某一个
// @Tags     ${Template.name}
// @Accept   json
// @Produce  json
// @Success  200  {object}  result.Result{data=Template}  true
// @Router   /Template/:id [get]
func getOne(ctx *gin.Context) {
	// 获取 id
	id := ctx.Param("id")

	// 获取当前 userId
	userId := token.GetUserId(ctx)

	// 根据 id、userId 查询
	one := template.FindByIdAndUserId(id, userId)

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
// @Router   /Template [get]
func getAll(ctx *gin.Context) {
	// 分页请求对象
	var req request.PageRequest

	// 获取当前 userId
	userId := token.GetUserId(ctx)

	// 较验参数
	app.ValidateParameter(ctx, &req)

	// 根据 userId 查询所有
	array, pagination := template.FindAllByUserId(userId, req.Page, req.PageSize)

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

// @Summary  创建
// @Tags     ${Template.name}
// @Accept   json
// @Produce  json
// @Param    json  body      postTemplateReq  true  "struct"
// @Success  200   {object}  result.Result    true
// @Router   /Template [post]
func postOne(ctx *gin.Context) {
	// 分页请求对象
	var req postTemplateReq

	// 较验参数
	app.ValidateParameter(ctx, &req)

	// 获取当前 userId
	userId := token.GetUserId(ctx)

	// 保存
	one := template.Template{
		UserId: userId,
		Name:   req.Name,
	}
	one.Save()

	// 返回结果
	response.New(ctx).IdData(one.Id)
}

// @Summary  更新
// @Tags     ${Template.name}
// @Accept   json
// @Produce  json
// @Param    json  body      putTemplateReq  true  "struct"
// @Success  200   {object}  result.Result   true
// @Router   /Template/:id [put]
func putOne(ctx *gin.Context) {
	// 分页请求对象
	var req putTemplateReq

	// 较验参数
	app.ValidateParameter(ctx, &req)

	// 获取当前 userId
	userId := token.GetUserId(ctx)

	// 更新
	one := template.Template{
		Name: req.Name,
	}
	template.UpdateByIdAndUserId(one.Id, userId, one)

	// 返回结果
	response.New(ctx).Ok()
}

// @Summary  删除
// @Tags     ${Template.name}
// @Accept   json
// @Produce  json
// @Success  200  {object}  result.Result  true
// @Router   /Template/ [put]
func deleteOne(ctx *gin.Context) {
	// 获取 id
	id := ctx.Param("id")

	// 获取当前 userId
	userId := token.GetUserId(ctx)

	// 删除
	template.DeleteByIdAndUserId(id, userId)

	// 返回结果
	response.New(ctx).Ok()
}
