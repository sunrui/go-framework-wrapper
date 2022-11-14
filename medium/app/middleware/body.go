/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-04-25 22:13:05
 */

package middleware

import (
	"github.com/gin-gonic/gin"
	"medium/app/request"
)

// Body 中间件
func Body(ctx *gin.Context) {
	if request.IsEnable() {
		request.CopyBody(ctx)
	}

	ctx.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	ctx.Next()
}
