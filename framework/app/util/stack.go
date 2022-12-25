/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-10-15 09:02:47
 */

package util

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

// StackFile 获取推栈层级
func StackFile(skip int, level int) []string {
	stacks := make([]string, 0)

	// 最大函数层级
	pc := make([]uintptr, level)
	runtime.Callers(skip, pc)
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
