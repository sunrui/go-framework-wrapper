/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-04-25 22:13:05
 */

package middleware

import (
	"fmt"
	"framework/result"
	"framework/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Recover 异常捕获中间件
func Recover(ctx *gin.Context) {
	// 捕获对象，全部抛出可以使用 panic 方法。
	defer func() {
		if err := recover(); err != nil {
			r := result.InternalError.WithData(result.M{
				"stack": util.Stack(5),
				"error": fmt.Sprintf("%s", err),
			})

			ctx.AbortWithStatusJSON(http.StatusOK, r)
		} else {
			ctx.Next()
		}
	}()
}
