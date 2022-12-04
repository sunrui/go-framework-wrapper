/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-12-04 15:10:15
 */

package glog

import (
	"fmt"
)

// Level 级别
type Level string

// String 字符串
func (level Level) String() string {
	return string(level)
}

const (
	Debug Level = "Debug" // 调试
	Info  Level = "Info"  // 信息
	Warn  Level = "Warn"  // 警告
	Error Level = "Error" // 错误
)

// GLog 日志
type GLog struct {
	Layout    Layout
	Appenders []Appender
}

// Write 写入
func (gLog GLog) Write(p []byte) (n int, err error) {
	for _, appender := range gLog.Appenders {
		appender.Print(Debug, gLog.Layout.getLayout(Debug, string(p)))
	}

	return len(p), nil
}

// Print 打印
func (gLog GLog) Print(level Level, format string, v ...interface{}) {
	for _, appender := range gLog.Appenders {
		message := fmt.Sprintf(format, v...)
		appender.Print(level, gLog.Layout.getLayout(level, message))
	}
}

// Println 打印并换行
func (gLog GLog) Println(level Level, format string, v ...interface{}) {
	for _, appender := range gLog.Appenders {
		message := fmt.Sprintf(format, v...)
		appender.Print(level, gLog.Layout.getLayout(level, message)+"\n")
	}
}

// Debug 调试
func (gLog GLog) Debug(format string, v ...interface{}) {
	gLog.Println(Debug, format, v...)
}

// Info 信息
func (gLog GLog) Info(format string, v ...interface{}) {
	gLog.Println(Info, format, v...)
}

// Warn 警告
func (gLog GLog) Warn(format string, v ...interface{}) {
	gLog.Println(Warn, format, v...)
}

// Error 错误
func (gLog GLog) Error(format string, v ...interface{}) {
	gLog.Println(Error, format, v...)
}
