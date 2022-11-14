/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-04-25 22:13:05
 */

package middleware

import (
	"github.com/gin-gonic/gin"
)

// MethodNotAllowed 异常 405 中间件
func MethodNotAllowed(ctx *gin.Context) {
	// 返回客户端
	//ctx.JSON(http.StatusOK, result.MethodNotAllowed.WithKeyPair("uri", ctx.Request.URL.RequestURI()))
	ctx.Abort()
}
