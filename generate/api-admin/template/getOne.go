/*
 * Copyright (c) $today.year honeysense.com All rights reserved.
 * Author: $author
 * Date: $today.format("yyyy-MM-dd HH:mm:ss")
 */

package template

import (
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
// @Failure  400  {object}  result.Result                 true  "{"code":"NoContent","message":"没有数据"}"
// @Router   /api-admin/template/:id [get]
func getOne(ctx *gin.Context) result.Result {
	// 获取 id
	id := ctx.Param("id")

	// 根据 id 查询
	if one := template.FindById(id); one == nil {
		return result.NoContent
	} else {
		return result.Ok
	}
}
