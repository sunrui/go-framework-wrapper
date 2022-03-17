/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/07 01:55:07
 */

package app_gin

import (
	"github.com/gin-gonic/gin"
	"medium-server-go/framework/result"
	"net/http"
)

// 统一返回对象
func Response(ctx *gin.Context, result result.Result) {
	ctx.JSON(http.StatusOK, result)
	ctx.Abort()
}
