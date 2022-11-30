/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-04-25 22:13:05
 */

package middleware

import (
	"fmt"
	"framework/http/response"
	"framework/result"
	"framework/util"
	"github.com/gin-gonic/gin"
)

// Recover 异常捕获中间件
func Recover() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				var r *result.Result
				var ok bool

				// 判断是否抛出了 Result
				if r, ok = err.(*result.Result); !ok {
					r = result.InternalError.WithData(result.M{
						"stack": util.Stack(10),
						"error": fmt.Sprintf("%s", err),
					})
				}

				response.Response(ctx, r)
			}
		}()

		ctx.Next()
	}
}
