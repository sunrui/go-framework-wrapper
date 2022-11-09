/*
 * Copyright (c) $today.year honeysense.com All rights reserved.
 * Author: $author
 * Date: $today.format("yyyy-MM-dd HH:mm:ss")
 */

package template

import (
	"framework/db"
	"framework/result"
	"framework/token"
	"framework/util"
	"generate/service/template"
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
// @RouterGroup   /api-user/template [post]
func postOne(ctx *gin.Context) result.Result {
	// 创建请求对象
	var req postOneReq

	// 较验参数
	util.ValidateParameter(ctx, &req)

	// 获取当前 userId
	userId := token.MustGetUserId(ctx)

	// 生成新对象
	one := template.Template{
		UserId: userId,
		Name:   req.Name,
	}

	// 保存
	if tx := db.Mysql.Save(one); tx.Error != nil {
		panic(tx.Error.Error())
	}

	// 返回结果
	return result.Ok.WithIdData(one.Id)
}
