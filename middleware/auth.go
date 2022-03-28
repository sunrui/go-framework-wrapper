/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author:sunrui
 * Date: 2022/01/03 07:51:03
 */

package middleware

import (
	"github.com/gin-gonic/gin"
	"medium-server-go/framework/app"
	"medium-server-go/framework/result"
	"medium-server-go/framework/token"
)

// AuthMiddleware 授权中间件
func AuthMiddleware(ctx *gin.Context) {
	_, err := token.Get(ctx)
	if err != nil {
		app.Result(ctx, result.NoAuth)
	}
}
