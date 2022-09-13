/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-04-25 20:47:56
 */

package middleware

import (
	"framework/proto/result"
	"framework/proto/token"
	"github.com/gin-gonic/gin"
	"service/user"
)

// Auth 授权中间件
func Auth(ctx *gin.Context) {
	if t, err := token.Get(ctx); err != nil {
		panic(result.NoAuth)
	} else {
		if user.FindById(t.UserId) == nil {
			panic(result.NoAuth)
		}
	}
}
