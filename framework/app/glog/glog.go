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
	DebugLevel Level = "Debug" // 调试
	InfoLevel  Level = "Info"  // 信息
	WarnLevel  Level = "Warn"  // 警告
	ErrorLevel Level = "Error" // 错误
)

// GLog 日志
type GLog struct {
	Layout    Layout     // 布局
	Appenders []Appender // 附加者
}

// NewGLog 创建
func NewGLog(layout Layout, appenders []Appender) *GLog {
	return &GLog{
		Layout:    layout,
		Appenders: appenders,
	}
}

// Write 写入
func (gLog GLog) Write(p []byte) (n int, err error) {
	for _, appender := range gLog.Appenders {
		appender.Print(DebugLevel, gLog.Layout.getLayout(DebugLevel, string(p)))
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

// PrintHttp 打印并换行
func (gLog GLog) PrintHttp(level Level, http Http) {
	for _, appender := range gLog.Appenders {
		appender.PrintHttp(level, http)
	}
}

// Debug 调试
func (gLog GLog) Debug(format string, v ...interface{}) {
	gLog.Println(DebugLevel, format, v...)
}

// Info 信息
func (gLog GLog) Info(format string, v ...interface{}) {
	gLog.Println(InfoLevel, format, v...)
}

// Warn 警告
func (gLog GLog) Warn(format string, v ...interface{}) {
	gLog.Println(WarnLevel, format, v...)
}

// Error 错误
func (gLog GLog) Error(format string, v ...interface{}) {
	gLog.Println(ErrorLevel, format, v...)
}
