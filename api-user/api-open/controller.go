/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/31 21:00:31
 */

package api_open

import (
	"github.com/gin-gonic/gin"
	"medium-server-go/common/app-gin"
	"medium-server-go/common/result"
	"medium-server-go/service/biz/open"
	"medium-server-go/service/provider"
)

// 获取指定用户下所有入驻
func getOpen(ctx *gin.Context) {
	// 获取当前用户 id
	userId := provider.Token.GetUserId(ctx)

	// 获取当前用户下的入驻
	opens := open.GetOpens(userId)

	app_gin.Response(ctx, result.Ok.WithData(opens))
}

// 提交入驻
func postOpen(ctx *gin.Context) {
	var req postOpenReq

	// 较验参数
	errData, err := app_gin.ValidateParameter(ctx, &req)
	if err != nil {
		app_gin.Response(ctx, result.ParameterError.WithData(errData))
		return
	}
}
