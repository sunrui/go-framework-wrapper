/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-12-04 15:10:15
 */

package glog

import (
	"fmt"
	"framework/app/result"
	"framework/server/request"
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

// Format 格式化
type Format struct {
	Level   Level            // 日志级别
	Message string           // 消息
	Request *request.Request // 请求
	Result  *result.Result   // 结果
	Elapsed int64            // 耗时
	UserId  *string          // 用户 id
}

// GLog 日志
type GLog struct {
	Layout    Layout
	Appenders []Appender
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

// PrintMessage 打印并换行
func (gLog GLog) PrintMessage(format Format) {
	for _, appender := range gLog.Appenders {
		appender.PrintMessage(format)
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
