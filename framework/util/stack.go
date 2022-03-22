/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022/03/17 22:44:17
 */

package util

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

// Stack 堆栈对象
type Stack struct {
	Function string // 函数
	File     string // 行数
}

// GetStack 获取推栈层级
func GetStack(deep int) []Stack {
	var stacks []Stack

	// 最大函数层级
	pc := make([]uintptr, deep)
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

	return stacks
}
