/*
 * Copyright (c) $today.year honeysense.com All rights reserved.
 * Author: $author
 * Date: $today.format("yyyy-MM-dd HH:mm:ss")
 */

package template

import (
	"framework/proto/response"
	"framework/proto/result"
	"framework/proto/token"
	"generate/service/core/template"
	"github.com/gin-gonic/gin"
)

// @Summary  删除
// @Tags     模板
// @Accept   json
// @Produce  json
// @Param    id   path      string         true  "id"
// @Success  200  {object}  result.Result  true
// @Failure  400  {object}  result.Result  true  "{"code":"NotFound","message":"不存在"}"
// @Router   /api-user/template/ [put]
func deleteOne(ctx *gin.Context) {
	// 获取 id
	id := ctx.Param("id")

	// 获取当前 userId
	userId := token.GetUserId(ctx)

	// 删除
	success := template.DeleteByIdAndUserId(id, userId)
	if !success {
		response.New(ctx).Result(result.NotFound.WithIdData(id))
		return
	}

	// 返回结果
	response.New(ctx).Ok()
}