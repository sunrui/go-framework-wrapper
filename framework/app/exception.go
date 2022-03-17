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
	"runtime"
	"strings"
)

// 异常捕获对象
func exceptionHandler(handlerFunc gin.HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 为了更好的调试，在开发环境中输出系统错误。
		if !config.IsDebugMode() {
			// 捕获对象，全部抛出可以使用 panic 方法。
			defer func() {
				if err := recover(); err != nil {
					dataMap := make(map[string]interface{})

					// 判断是否抛出了 result 对象
					res, ok := err.(*result.Result)
					if ok {
						dataMap["error"] = res.Data
					} else {
						dataMap["error"] = err
					}

					mapData := make(map[string]interface{})
					mapData["description"] = err

					// 栈堆对象
					type Stack struct {
						Function string // 函数
						File     string // 行数
					}

					var stacks []Stack

					// 最大函数位深 5 层
					maxDeep := 6
					pc := make([]uintptr, maxDeep)
					runtime.Callers(4, pc)
					frames := runtime.CallersFrames(pc)

					// 当前项目目录
					pwd, _ := os.Getwd()
					pwd = strings.Replace(pwd, "\\", "/", -1)

					// 当前 go 目录
					goPath := os.Getenv("GOPATH")
					goPath = strings.Replace(goPath, "\\", "/", -1)

					for frame, ok := frames.Next(); ok; frame, ok = frames.Next() {
						// 去掉项目目录
						file := strings.Replace(frame.File, pwd, "", -1)
						file = strings.Replace(file, goPath, "", -1)
						file = fmt.Sprintf("%s:%d", file, frame.Line)
						function := frame.Function[strings.Index(frame.Function, "/"):]

						stacks = append(stacks, Stack{
							Function: function,
							File:     file,
						})
					}

					dataMap["stacks"] = stacks

					Response(ctx, result.InternalError.WithData(dataMap))
				}
			}()
		}

		ctx.Next()
		handlerFunc(ctx)
	}
}
