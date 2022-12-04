/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-04-25 22:13:05
 */

package middleware

import (
	"framework/app/result"
	"github.com/gin-gonic/gin"
)

// NotFound 异常 404 中间件
func NotFound(ctx *gin.Context) *result.Result {
	ctx.Abort()

	return result.NotFound.WithData(result.M{
		"uri": ctx.Request.URL.RequestURI(),
	})
}
