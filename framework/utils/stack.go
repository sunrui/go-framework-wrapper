/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022/04/16 17:02:16
 */

package utils

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

// DumpStack 获取推栈层级
func DumpStack(level int) []string {
	var stacks = make([]string, 0)

	// 最大函数层级
	pc := make([]uintptr, level)
	runtime.Callers(3, pc)
	frames := runtime.CallersFrames(pc)

	// 当前项目目录
	pwd, _ := os.Getwd()
	pwd = pwd[:strings.LastIndex(pwd, "/")]

	for frame, ok := frames.Next(); ok; frame, ok = frames.Next() {
		// 过滤掉系统目录
		if !strings.HasPrefix(frame.File, pwd) {
			continue
		}

		// 去掉项目目录
		file := strings.Replace(frame.File, pwd, "", -1)
		file = fmt.Sprintf("%s:%d", file, frame.Line)
		stacks = append(stacks, file)
	}

	return stacks
}
