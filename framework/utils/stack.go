/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022/04/16 17:02:16
 */

package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// 堆栈对象
type stack struct {
	Function string `json:"function"` // 函数
	File     string `json:"file"`     // 行数
}

// GetStack 获取推栈层级
func GetStack(level int) *stack {
	// 最大函数层级
	pc := make([]uintptr, level)
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

		return &stack{
			Function: function,
			File:     file,
		}
	}

	return &stack{}
}
