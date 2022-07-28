/*
 * Copyright (c) $today.year honeysense.com All rights reserved.
 * Author: $author
 * Date: $today.format("yyyy-MM-dd HH:mm:ss")
 */

package template

import (
	"framework/app"
	"framework/proto/response"
	"framework/proto/token"
	"generate/service/core/template"
	"github.com/gin-gonic/gin"
)

// 创建请求
type postOneReq struct {
	Name string `json:"name" validate:"required"` // 名称
}

// @Summary  创建
// @Tags     模板
// @Accept   json
// @Produce  json
// @Param    json  body      postOneReq  true  "json"
// @Success  200   {object}  result.Result    true
// @Router   /api-user/template [post]
func postOne(ctx *gin.Context) {
	// 创建请求对象
	var req postOneReq

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
