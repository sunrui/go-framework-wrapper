/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-04-25 22:14:59
 */

package auth

import (
	"framework/proto/result"
	"framework/proto/token"
	"github.com/gin-gonic/gin"
)

// 登出
func postLogout(ctx *gin.Context) result.Result {
	_, err := ctx.Cookie("token")
	if err == nil {
		token.Remove(ctx)
	}

	return result.Ok
}
