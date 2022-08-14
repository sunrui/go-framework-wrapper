/*
 * Copyright (c) $today.year honeysense.com All rights reserved.
 * Author: $author
 * Date: $today.format("yyyy-MM-dd HH:mm:ss")
 */

package template

import (
	"framework/proto/result"
	"generate/service/template"
	"github.com/gin-gonic/gin"
)

// @Summary  删除
// @Tags     模板
// @Accept   json
// @Produce  json
// @Param    id   path      string         true  "id"
// @Success  200  {object}  result.Result  true
// @Failure  400  {object}  result.Result  true  "{"code":"NotFound","message":"不存在"}"
// @Router   /api-admin/template/ [put]
func deleteOne(ctx *gin.Context) result.Result {
	// 获取 id
	id := ctx.Param("id")

	// 删除
	if ok := template.DeleteById(id); !ok {
		return result.NotFound.WithIdData(id)
	} else {
		return result.Ok
	}
}
