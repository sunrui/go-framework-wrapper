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

// MethodNotAllowed 异常 405 中间件
func MethodNotAllowed(ctx *gin.Context) *result.Result[any] {
	ctx.Abort()

	return &result.Result[any]{
		Code: result.METHOD_NOT_ALLOWED,
		Data: result.KeyValueData("uri", ctx.Request.URL.RequestURI()),
	}
}
