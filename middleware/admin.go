/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/26 13:40:26
 */

package middleware

import (
	"github.com/gin-gonic/gin"
	"medium-server-go/framework/app"
	"medium-server-go/framework/result"
	"medium-server-go/framework/token"
)

// AdminMiddleware 管理中间件
func AdminMiddleware(ctx *gin.Context) {
	_, err := token.Get(ctx)
	if err != nil {
		app.Response(ctx, result.NoAuth)
	}
}
