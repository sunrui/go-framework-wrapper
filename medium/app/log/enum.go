/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-14 23:32:57
 */

package log

// Level 日志等级
type LevelType string

const (
	NONE    LevelType = "NONE"    // none
	TRACE   LevelType = "TRACE"   // trace
	DEBUG   LevelType = "DEBUG"   // debug
	INFO    LevelType = "INFO"    // info
	WARNING LevelType = "WARNING" // warning
	ERROR   LevelType = "ERROR"   // error
)
