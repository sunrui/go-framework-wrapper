/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-04-25 20:48:06
 */

package middleware

import (
	"framework/app/token"
	"github.com/gin-gonic/gin"
)

// Admin 管理中间件
func Admin(ctx *gin.Context) {
	_ = token.MustGetUserId(ctx)
}
