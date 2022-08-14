/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-04-25 22:14:59
 */

package auth

import (
	"framework/app"
	"framework/proto/result"
	"github.com/gin-gonic/gin"
)

// 微信登录
func postLoginByWechat(ctx *gin.Context) result.Result {
	var req postLoginByPhoneReq

	// 较验参数
	app.ValidateParameter(ctx, &req)

	return result.Ok
}
