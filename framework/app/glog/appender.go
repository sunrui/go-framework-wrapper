/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-12-04 17:23:12
 */

package glog

import (
	"fmt"
)

// Appender 附加器
type Appender interface {
	// Print 打印
	Print(level Level, message string)
	// PrintHttp 打印 http
	PrintHttp(level Level, http Http)
}

// ConsoleAppender 控制台附加器
type ConsoleAppender struct {
}

// Print 打印
func (consoleAppender ConsoleAppender) Print(level Level, message string) {
	switch level {
	case DebugLevel:
		message = fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", 30, message)
	case InfoLevel:
		message = fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", 32, message)
	case WarnLevel:
		message = fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", 33, message)
	case ErrorLevel:
		message = fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", 31, message)
	}

	fmt.Print(message)
}

// PrintHttp 打印 http
func (consoleAppender ConsoleAppender) PrintHttp(level Level, http Http) {
	consoleAppender.Print(level, http.LineString())
}

// FileAppender 文件附加器
type FileAppender struct {
	Debug *RotateLog // 调试
	Info  *RotateLog // 信息
	Warn  *RotateLog // 警告
	Error *RotateLog // 错误
}

// Print 打印
func (fileAppender FileAppender) Print(level Level, message string) {
	switch level {
	case DebugLevel:
		if fileAppender.Debug != nil {
			fileAppender.Debug.Debug(message)
		}
	case InfoLevel:
		if fileAppender.Info != nil {
			fileAppender.Info.Info(message)
		}
	case WarnLevel:
		if fileAppender.Warn != nil {
			fileAppender.Warn.Warn(message)
		}
	case ErrorLevel:
		if fileAppender.Error != nil {
			fileAppender.Error.Error(message)
		}
	}
}

// PrintHttp 打印 http
func (fileAppender FileAppender) PrintHttp(level Level, http Http) {
	fileAppender.Print(level, http.LineString())
}
