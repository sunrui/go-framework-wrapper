/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-04-25 22:13:05
 */

package middleware

import (
	"framework/config"
	"framework/request"
	"github.com/gin-gonic/gin"
)

// BodyMiddleware body 中间件
func BodyMiddleware(ctx *gin.Context) {
	// 如果需要记录日志或请求被异出则拷贝 body 对象
	if config.Get().Log.Enable || request.IsDebugRequest(ctx) {
		request.CopyBody(ctx)
	}

	ctx.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	ctx.Next()
}
