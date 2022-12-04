/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-04-25 22:13:05
 */

package middleware

import (
	"fmt"
	"framework/app/result"
	"framework/app/util"
	"github.com/gin-gonic/gin"
)

// Recover 异常捕获中间件
func Recover(ctx *gin.Context) (r *result.Result) {
	defer func() {
		if err := recover(); err != nil {
			var ok bool

			// 判断是否抛出了 Result
			if r, ok = err.(*result.Result); !ok {
				r = result.InternalError.WithData(result.M{
					"stack": util.Stack(10),
					"error": fmt.Sprintf("%s", err),
				})
			}
		}
	}()

	ctx.Next()

	return r
}
