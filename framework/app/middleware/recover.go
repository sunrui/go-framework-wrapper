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
)

// Recover 异常捕获中间件
func Recover(ctx *gin.Context) *result.Result {
	// 捕获对象，全部抛出可以使用 panic 方法。
	defer func() *result.Result {
		var r *result.Result = nil

		if err := recover(); err != nil {
			ctx.Abort()

			result.InternalError.WithData(result.M{
				"stack": util.Stack(5),
				"error": fmt.Sprintf("%s", err),
			})
		}

		return r
	}()

	ctx.Next()
	return nil
}
