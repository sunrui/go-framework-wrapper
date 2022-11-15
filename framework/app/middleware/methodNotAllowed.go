/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-04-25 22:13:05
 */

package middleware

import (
	"framework/result"
	"github.com/gin-gonic/gin"
)

// MethodNotAllowed 异常 405 中间件
func MethodNotAllowed(ctx *gin.Context) *result.Result[any] {
	ctx.Abort()

	return &result.Result[any]{
		Code: result.MethodNotAllowed,
		Data: result.KeyValueData("uri", ctx.Request.URL.RequestURI()),
	}
}
