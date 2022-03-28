/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/03 07:51:03
 */

package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"medium-server-go/framework/config"
	"medium-server-go/framework/result"
	"os"
	"path/filepath"
	"runtime"
	"runtime/debug"
	"strings"
)

// 堆栈对象
type stack struct {
	Function string `json:"function"` // 函数
	File     string `json:"file"`     // 行数
}

// 获取推栈层级
func getStack() stack {
	// 最大函数层级 5
	pc := make([]uintptr, 5)
	runtime.Callers(3, pc)
	frames := runtime.CallersFrames(pc)

	// 当前项目目录
	pwd, _ := os.Getwd()

	for frame, ok := frames.Next(); ok; frame, ok = frames.Next() {
		// 过滤掉系统目录
		if !strings.HasPrefix(frame.File, pwd) {
			continue
		}

		// 去掉项目目录
		file := strings.Replace(frame.File, pwd, "", -1)
		file = fmt.Sprintf("%s:%d", file, frame.Line)
		function := filepath.Base(frame.Function)

		return stack{
			Function: function,
			File:     file,
		}
	}

	return stack{}
}

// 异常捕获对象
func exceptionHandler(handlerFunc gin.HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 捕获对象，全部抛出可以使用 panic 方法。
		defer func() {
			if err := recover(); err != nil {
				dataMap := make(map[string]interface{})
				dataMap["stack"] = getStack()

				// 判断是否抛出了 result 对象
				res, ok := err.(*result.Result)
				if ok {
					dataMap["error"] = res.Data
				} else {
					dataMap["error"] = err
				}

				Error(ctx, result.InternalError.WithData(dataMap))

				// 为了更好的调试，在开发环境中输出系统错误。
				if config.IsDebugMode() {
					debug.PrintStack()
				}
			} else {
				ctx.Next()
			}
		}()

		handlerFunc(ctx)
	}
}
