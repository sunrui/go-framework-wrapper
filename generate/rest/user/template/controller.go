/*
 * Copyright (c) $today.year honeysense.com All rights reserved.
 * Author: sunrui
 * Date: $today.format("yyyy-MM-dd HH:mm:ss")
 */

package template

import (
	"framework/app"
	"framework/proto/request"
	"framework/proto/response"
	"framework/proto/result"
	"framework/proto/token"
	"generate/service/template"
	"github.com/gin-gonic/gin"
)

// @Summary  获取某一个
// @Tags     模板
// @Accept   json
// @Produce  json
// @Param    id   path      string                        true  "id"
// @Success  200  {object}  result.Result{data=Template}  true
// @Failure  400  {object}  result.Result                 true  "{"code":"NoData","message":"没有数据"}"
// @Router   /user/Template/:id [get]
func getOne(ctx *gin.Context) {
	// 获取 id
	id := ctx.Param("id")

	// 获取当前 userId
	userId := token.GetUserId(ctx)

	// 根据 id、userId 查询
	one := template.FindByIdAndUserId(id, userId)

	// 未找到结果
	if one == nil {
		response.New(ctx).Data(result.NoData)
		return
	}

	// 返回结果
	response.New(ctx).Data(one)
}

// @Summary  获取所有
// @Tags     模板
// @Accept   json
// @Produce  json
// @Param    page      query     int                                 true  "分页，从 1 开始"
// @Param    pageSize  query     int                                 true  "分页大小"
// @Success  200       {object}  result.PageResult{data=[]Template}  true
// @Failure  400       {object}  result.Result                       true  "{"code":"NoData","message":"没有数据"}"
// @Router   /user/Template [get]
func getAll(ctx *gin.Context) {
	// 分页请求对象
	var req request.PageRequest

	// 获取当前 userId
	userId := token.GetUserId(ctx)

	// 较验参数
	app.ValidateParameter(ctx, &req)

	// 根据 userId 查询所有
	array, pagination := template.FindAllByUserId(userId, req.Page, req.PageSize, true)

	// 未找到结果
	if len(array) == 0 {
		response.New(ctx).Data(result.NoData)
		return
	}

	// 返回结果
	response.New(ctx).Data(result.PageResult{
		Result:     result.Ok.WithData(array),
		Pagination: pagination,
	})
}

// @Summary  创建
// @Tags     模板
// @Accept   json
// @Produce  json
// @Param    json  body      postTemplateReq  true  "json"
// @Success  200   {object}  result.Result    true
// @Router   /user/Template [post]
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
// @Tags     模板
// @Accept   json
// @Produce  json
// @Param    id    path      string          true  "id"
// @Param    json  body      putTemplateReq  true  "json"
// @Success  200   {object}  result.Result   true
// @Failure  400   {object}  result.Result   true  "{"code":"NotFound","message":"不存在"}"
// @Router   /user/Template/:id [put]
func putOne(ctx *gin.Context) {
	// 分页请求对象
	var req putTemplateReq

	// 较验参数
	app.ValidateParameter(ctx, &req)

	// 获取 id
	id := ctx.Param("id")

	// 获取当前 userId
	userId := token.GetUserId(ctx)

	// 更新
	one := template.Template{
		Name: req.Name,
	}

	success := template.UpdateByIdAndUserId(id, userId, one)
	if !success {
		response.New(ctx).Data(result.NotFound.WithIdData(id))
		return
	}

	// 返回结果
	response.New(ctx).Ok()
}

// @Summary  删除
// @Tags     模板
// @Accept   json
// @Produce  json
// @Param    id   path      string         true  "id"
// @Success  200  {object}  result.Result  true
// @Failure  400  {object}  result.Result  true  "{"code":"NotFound","message":"不存在"}"
// @Router   /user/Template/ [put]
func deleteOne(ctx *gin.Context) {
	// 获取 id
	id := ctx.Param("id")

	// 获取当前 userId
	userId := token.GetUserId(ctx)

	// 删除
	success := template.DeleteByIdAndUserId(id, userId)
	if !success {
		response.New(ctx).Data(result.NotFound.WithIdData(id))
		return
	}

	// 返回结果
	response.New(ctx).Ok()
}
