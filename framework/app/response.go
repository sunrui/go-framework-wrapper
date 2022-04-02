/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/07 01:55:07
 */

package app

import (
	"github.com/gin-gonic/gin"
	"medium-server-go/framework/result"
	"net/http"
)

// 结果对象实体
type responseDef struct {
	ctx *gin.Context
}

// Response 结果对象
func Response(ctx *gin.Context) *responseDef {
	return &responseDef{
		ctx: ctx,
	}
}

// Ok 操作成功返回对象
func (responseDef *responseDef) Ok() {
	responseDef.ctx.JSON(http.StatusOK, result.Ok)
	responseDef.ctx.Abort()
}

// Data 数据返回对象
func (responseDef *responseDef) Data(data interface{}) {
	responseDef.ctx.JSON(http.StatusOK, result.Ok.WithData(data))
	responseDef.ctx.Abort()
}

// IdData 主键返回对象
func (responseDef *responseDef) IdData(id string) {
	type idData struct {
		Id string `json:"id"`
	}

	responseDef.Data(idData{
		Id: id,
	})
}

// PageData 分页数据返回对象
func (responseDef *responseDef) PageData(data interface{}, pagination result.Pagination) {
	responseDef.ctx.JSON(http.StatusOK, result.PageResult{
		Result:     result.Ok.WithData(data),
		Pagination: pagination,
	})
	responseDef.ctx.Abort()
}
