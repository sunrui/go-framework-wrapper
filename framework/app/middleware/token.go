/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-04-25 22:13:05
 */

package middleware

import (
	"github.com/gin-gonic/gin"
)

// Token 令牌中间件
func Token(ctx *gin.Context) {
	ctx.Next()
}
