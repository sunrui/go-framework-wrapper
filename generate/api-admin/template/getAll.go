/*
 * Copyright (c) $today.year honeysense.com All rights reserved.
 * Author: $author
 * Date: $today.format("yyyy-MM-dd HH:mm:ss")
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

// @Summary  获取所有
// @Tags     模板
// @Accept   json
// @Produce  json
// @Param    page      query     int                                 true  "分页，从 1 开始"
// @Param    pageSize  query     int                                 true  "分页大小"
// @Success  200       {object}  result.PageResult{data=[]Template}  true
// @Failure  400       {object}  result.Result                       true  "{"code":"NoContent","message":"没有数据"}"
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
		response.New(ctx).Result(result.NoContent)
		return
	}

	// 返回结果
	response.New(ctx).PageResult(result.PageResult{
		Result:     result.Ok.WithData(array),
		Pagination: pagination,
	})
}
