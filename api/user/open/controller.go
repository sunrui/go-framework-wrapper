/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/31 21:00:31
 */

package open

import (
	"github.com/gin-gonic/gin"
	"medium-server-go/framework/app"
	"medium-server-go/framework/token"
	"medium-server-go/service/open/open"
)

// 获取指定用户下所有入驻
func getOpen(ctx *gin.Context) {
	// 获取当前用户 id
	userId := token.GetUserId(ctx)

	// 获取当前用户下的入驻
	opens := open.GetOpen(userId)
	app.Result(ctx).Data(opens)
}

// 提交入驻
func postOpen(ctx *gin.Context) {
	var req postOpenReq

	// 较验参数
	app.ValidateParameter(ctx, &req)
}
