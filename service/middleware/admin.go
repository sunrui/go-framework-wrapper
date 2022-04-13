/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022/04/07 22:33:07
 */

package middleware

import (
	"framework/proto/response"
	"framework/proto/result"
	"framework/proto/token"
	"github.com/gin-gonic/gin"
)

// Admin 管理中间件
func Admin(ctx *gin.Context) {
	if _, err := token.Get(ctx); err != nil {
		response.New(ctx).Data(result.NoAuth)
	}
}
