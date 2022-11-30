/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-04-25 20:47:56
 */

package middleware

import (
	"framework/context"
	"github.com/gin-gonic/gin"
)

// Auth 授权中间件
func Auth(ctx *gin.Context) {
	_ = context.Token.MustGetUserId(ctx)
}
