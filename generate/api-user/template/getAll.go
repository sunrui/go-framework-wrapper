/*
 * Copyright (c) $today.year honeysense.com All rights reserved.
 * Author: $author
 * Date: $today.format("yyyy-MM-dd HH:mm:ss")
 */

package template

import (
	"framework/app"
	"framework/proto/request"
	"framework/proto/result"
	"framework/proto/token"
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
// @RouterGroup   /api-user/template [get]
func getAll(ctx *gin.Context) result.Result {
	// 分页请求对象
	var req request.PageRequest

	// 获取当前 userId
	userId := token.MustGetUserId(ctx)

	// 较验参数
	app.ValidateParameter(ctx, &req)

	// 根据 userId 查询所有
	if array, pagination := template.FindAllByUserId(userId, req.Page, req.PageSize, true); len(array) == 0 {
		return result.NoContent
	} else {
		return result.Ok.WithPageData(array, pagination)
	}
}
