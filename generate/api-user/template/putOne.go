/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022/07/28 20:16:28
 */

package template

import (
	"framework/app"
	"framework/proto/result"
	"framework/proto/token"
	"generate/service/core/template"
	"github.com/gin-gonic/gin"
)

// 更新请求
type putOneReq struct {
	Name string `json:"name" validate:"required"` // 名称
}

// @Summary  更新
// @Tags     模板
// @Accept   json
// @Produce  json
// @Param    id    path      string          true  "id"
// @Param    json  body      putOneReq  true  "json"
// @Success  200   {object}  result.Result   true
// @Failure  400   {object}  result.Result   true  "{"code":"NotFound","message":"不存在"}"
// @RouterGroup   /api-user/template/:id [put]
func putOne(ctx *gin.Context) result.Result {
	// 更新请求对象
	var req putOneReq

	// 较验参数
	app.ValidateParameter(ctx, &req)

	// 获取 id
	id := ctx.Param("id")

	// 获取当前 userId
	userId := token.MustGetUserId(ctx)

	// 更新
	one := template.Template{
		Name: req.Name,
	}

	if ok := template.UpdateByIdAndUserId(id, userId, one); !ok {
		return result.NotFound.WithIdData(id)
	} else {
		return result.Ok
	}
}
