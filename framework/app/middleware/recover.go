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
			// 判断是否抛出了 Result 对象
			if res, ok := err.(result.Result); ok {
				ctx.JSON(http.StatusOK, res)
			} else {
				// 堆栈信息
				dataMap := make(map[string]any)
				dataMap["stack"] = util.Stack(5)
				dataMap["error"] = fmt.Sprintf("%s", err)

				ctx.JSON(http.StatusOK, result.InternalError.WithData(dataMap))
				ctx.Abort()
			}
		}
	}()

	ctx.Next()
}
