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
	if _, err := ctx.Cookie("token"); err != nil {
		return result.NoAuth
	} else {
		token.Remove(ctx)
		return result.Ok
	}
}
