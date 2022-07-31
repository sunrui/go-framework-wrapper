/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-04-25 22:14:59
 */

package auth

import (
	"framework/proto/response"
	"framework/proto/result"
	"framework/proto/token"
	"github.com/gin-gonic/gin"
)

// 登出
func postLogout(ctx *gin.Context) {
	_, err := ctx.Cookie("token")
	if err != nil {
		response.Result(ctx, result.NotFound.WithData(err.Error()))
		return
	}

	// 移除令牌
	token.Remove(ctx)

	response.Result(ctx, result.Ok)
}
