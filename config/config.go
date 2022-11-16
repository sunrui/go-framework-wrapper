/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-07 00:15:14
 */

package config

import "github.com/sirupsen/logrus"

// Mysql 数据库
type Mysql struct {
	User     string `json:"user"`     // 用户名
	Password string `json:"password"` // 密码
	Host     string `json:"host"`     // 主机
	Port     int    `json:"port"`     // 端口
	Database string `json:"database"` // 数据库
}

// RateLimit 限流
type RateLimit struct {
	Capacity int64 `json:"capacity"` // 令牌桶容量
	Quantum  int64 `json:"quantum"`  // 每隔多少秒
}

// Redis 缓存
type Redis struct {
	Host     string `json:"host"`     // 主机
	Port     int    `json:"port"`     // 端口
	Password string `json:"password"` // 密码
	Database int    `json:"database"` // 数据库
}

// Log 日志
type Log struct {
	Directory string       `json:"directory"` // 路径
	File      string       `json:"file"`      // 文件
	Level     logrus.Level `json:"level"`     // 等级
}

// Request 请求
type Request struct {
	Dump bool `json:"dump"` // 导出
}

// Token 令牌
type Token struct {
	JwtSecret      string `json:"jwtSecret"`  // jwt 密钥
	Key            string `json:"key"`        // 键名
	MaxAge         int64  `json:"max_age"`    // 过期时间（秒）
	AutoRefreshAge int64  `json:"refreshAge"` // 自动重新刷新时间（秒）
}
