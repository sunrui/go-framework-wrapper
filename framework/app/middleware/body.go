/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-04-25 22:13:05
 */

package middleware

import (
	"framework/app/request"
	"github.com/gin-gonic/gin"
)

// Body 中间件
func Body(ctx *gin.Context) {
	// 将 body 数据缓存，用于返回给 request 和记录日志。
	request.CopyBody(ctx)

	ctx.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	ctx.Next()
}
