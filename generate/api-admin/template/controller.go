/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022/04/16 15:25:16
 */

package template

import (
	"framework/app"
	"framework/proto/request"
	"framework/proto/response"
	"framework/proto/result"
	"generate/service/core/template"
	"github.com/gin-gonic/gin"
)

// @Summary  获取某一个
// @Tags     模板
// @Accept   json
// @Produce  json
// @Param    id   path      string                        true  "id"
// @Success  200  {object}  result.Result{data=Template}  true
// @Failure  400  {object}  result.Result                 true  "{"code":"NoData","message":"没有数据"}"
// @Router   /api-admin/template/:id [get]
func getOne(ctx *gin.Context) {
	// 获取 id
	id := ctx.Param("id")

	// 根据 id 查询
	one := template.FindById(id)

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
// @Router   /api-admin/template [get]
func getAll(ctx *gin.Context) {
	// 分页请求对象
	var req request.PageRequest

	// 较验参数
	app.ValidateParameter(ctx, &req)

	// 根据 userId 查询所有
	array, pagination := template.FindAll(req.Page, req.PageSize, true)

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
// @Router   /api-admin/template [post]
func postOne(ctx *gin.Context) {
	// 创建请求对象
	var req postTemplateReq

	// 较验参数
	app.ValidateParameter(ctx, &req)

	// 获取当前 userId
	userId := "token.GetUserId(ctx)"

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
// @Router   /api-admin/template/:id [put]
func putOne(ctx *gin.Context) {
	// 更新请求对象
	var req putTemplateReq

	// 较验参数
	app.ValidateParameter(ctx, &req)

	// 获取 id
	id := ctx.Param("id")

	// 更新
	one := template.Template{
		Name: req.Name,
	}

	// 更新
	success := template.UpdateById(id, one)
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
// @Router   /api-admin/template/ [put]
func deleteOne(ctx *gin.Context) {
	// 获取 id
	id := ctx.Param("id")

	// 删除
	success := template.DeleteById(id)
	if !success {
		response.New(ctx).Data(result.NotFound.WithIdData(id))
		return
	}

	// 返回结果
	response.New(ctx).Ok()
}
