/*
 * Copyright (c) $today.year honeysense.com All rights reserved.
 * Author: $author
 * Date: $today.format("yyyy-MM-dd HH:mm:ss")
 */

package template

import (
	"framework/result"
	"framework/token"
	"generate/service/template"
	"github.com/gin-gonic/gin"
)

// @Summary  获取某一个
// @Tags     模板
// @Accept   json
// @Produce  json
// @Param    id   path      string                        true  "id"
// @Success  200  {object}  result.Result{data=Template}  true
// @Failure  400  {object}  result.Result                 true  "{"code":"NoContent","message":"没有数据"}"
// @Router   /api-user/template/:id [get]
func getOne(ctx *gin.Context) result.Result {
	// 获取 id
	id := ctx.Param("id")

	// 获取当前 userId
	userId := token.MustGetUserId(ctx)

	// 根据 id、userId 查询
	if one := template.FindByIdAndUserId(id, userId); one == nil {
		return result.NoContent
	} else {
		return result.Ok.WithData(one)
	}
}
