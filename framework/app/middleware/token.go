/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-04-25 22:13:05
 */

package middleware

import (
	"framework/app/token"
	"github.com/gin-gonic/gin"
)

// Token 注册刷新令牌中间件
func Token(ctx *gin.Context) {
	token.RefreshIf(ctx)
	ctx.Next()
}
