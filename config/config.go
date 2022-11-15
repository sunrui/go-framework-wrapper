/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-07 00:15:14
 */

package config

type Mysql struct {
	User     string `json:"user"`     // 用户名
	Password string `json:"password"` // 密码
	Host     string `json:"host"`     // 主机
	Port     int    `json:"port"`     // 端口
	Database string `json:"database"` // 数据库
}

type RateLimit struct {
	Capacity int64 `json:"capacity"` // 令牌桶容量
	Quantum  int64 `json:"quantum"`  // 每隔多少秒
}

type Redis struct {
	Host     string `json:"host"`     // 主机
	Port     int    `json:"port"`     // 端口
	Password string `json:"password"` // 密码
	Database int    `json:"database"` // 数据库
}

// LogLevel 日志等级
type LogLevel string

const (
	LogNone    LogLevel = "NONE"    // none
	LogTrace   LogLevel = "TRACE"   // trace
	LogDebug   LogLevel = "DEBUG"   // debug
	LogInfo    LogLevel = "INFO"    // info
	LogWarning LogLevel = "WARNING" // warning
	LogError   LogLevel = "ERROR"   // error
)

type Log struct {
	Directory string   `json:"directory"` // 路径
	Level     LogLevel `json:"level"`     // 等级
}

type Jwt struct {
	Secret string `json:"secret"` // 密钥
}

type Request struct {
	Dump bool `json:"dump"` // 导出
}
