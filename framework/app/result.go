/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/07 01:55:07
 */

package app

import (
	"github.com/gin-gonic/gin"
	"medium-server-go/framework/exception"
	"net/http"
)

// 结果对象实体
type resultDef struct {
	ctx *gin.Context
}

// 结果对象
func Result(ctx *gin.Context) *resultDef {
	return &resultDef{
		ctx: ctx,
	}
}

// Ok 操作成功返回对象
func (resultDef *resultDef) Ok() {
	resultDef.ctx.Status(http.StatusOK)
	resultDef.ctx.Abort()
}

// Data 数据返回对象
func (resultDef *resultDef) Data(data interface{}) {
	resultDef.ctx.JSON(http.StatusOK, data)
	resultDef.ctx.Abort()
}

// Id 主键返回对象
func (resultDef *resultDef) Id(id string) {
	type idData struct {
		Id string `json:"id"`
	}

	resultDef.Data(idData{
		Id: id,
	})
}

// Exception 异常返回对象
func (resultDef *resultDef) Exception(exception exception.Exception) {
	resultDef.ctx.JSON(http.StatusBadRequest, exception)
	resultDef.ctx.Abort()
}
