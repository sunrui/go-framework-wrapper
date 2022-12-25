/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-10-15 09:02:47
 */

package util

import (
	"fmt"
	"runtime"
)

// StackFile 获取推栈层级
func StackFile(skip int, level int) []string {
	stacks := make([]string, 0)

	pc := make([]uintptr, level)
	runtime.Callers(skip, pc)
	frames := runtime.CallersFrames(pc)

	for frame, ok := frames.Next(); ok; frame, ok = frames.Next() {
		stacks = append(stacks, fmt.Sprintf("%s:%d", frame.File, frame.Line))
	}

	return stacks
}
