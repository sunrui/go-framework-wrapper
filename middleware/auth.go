/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author:sunrui
 * Date: 2022/01/03 07:51:03
 */

package middleware

import (
	"github.com/gin-gonic/gin"
	"medium-server-go/framework/proto/response"
	"medium-server-go/framework/proto/result"
	"medium-server-go/framework/proto/token"
)

// Auth 授权中间件
func Auth(ctx *gin.Context) {
	_, err := token.Get(ctx)
	if err != nil {
		response.New(ctx).Data(result.NoAuth)
	}
}
