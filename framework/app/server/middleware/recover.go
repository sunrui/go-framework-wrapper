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
			if ret, ok := err.(result.Result); ok {
				r = &ret
			} else {
				ret = result.InternalError.WithData(result.M{
					"error": fmt.Sprintf("%s", err),
					"file":  util.StackFile(0, 10),
				})

				r = &ret
			}
		}
	}()

	ctx.Next()

	return r
}
