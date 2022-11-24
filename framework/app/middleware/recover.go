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
func Recover() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				r := result.InternalError.WithData(result.M{
					"stack": util.Stack(10),
					"error": fmt.Sprintf("%s", err),
				})

				ctx.AbortWithStatusJSON(http.StatusOK, r)
			}
		}()
		ctx.Next()
	}
}
