/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-04-25 22:13:05
 */

package middleware

import (
	"github.com/gin-gonic/gin"
)

// NotFound 异常 404 中间件
func NotFound(ctx *gin.Context) {
	//ctx.JSON(http.StatusOK, result.NotFound.WithKeyPair("uri", ctx.Request.URL.RequestURI()))
	ctx.Abort()
}
