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
func Recover(ctx *gin.Context) *result.Result[any] {
	// 捕获对象，全部抛出可以使用 panic 方法。
	defer func() *result.Result[any] {
		var r *result.Result[any] = nil

		if err := recover(); err != nil {
			// 堆栈信息
			dataMap := make(map[string]any)
			dataMap["stack"] = util.Stack(5)
			dataMap["error"] = fmt.Sprintf("%s", err)
			ctx.Abort()

			r = &result.Result[any]{
				Code: result.InternalError,
				Data: dataMap,
			}
		}

		return r
	}()

	ctx.Next()
	return nil
}
