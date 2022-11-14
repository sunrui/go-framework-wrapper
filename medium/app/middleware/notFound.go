/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-04-25 22:13:05
 */

package middleware

import (
	"github.com/gin-gonic/gin"
	"medium/result"
)

// NotFound 异常 404 中间件
func NotFound(ctx *gin.Context) *result.Result[any] {
	ctx.Abort()

	return &result.Result[any]{
		Code: result.NOT_FOUND,
		Data: result.KeyValueData("uri", ctx.Request.URL.RequestURI()),
	}
}
