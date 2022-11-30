/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-07 00:15:14
 */

package config

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"os"
)

// Mysql 数据库
type Mysql struct {
	User         string `json:"user"`         // 用户名
	Password     string `json:"password"`     // 密码
	Host         string `json:"host"`         // 主机
	Port         int    `json:"port"`         // 端口
	Database     string `json:"database"`     // 数据库
	MaxOpenConns int    `json:"maxOpenConns"` // 最大打开连接
	MaxIdleConns int    `json:"maxIdleConns"` // 最大空闲连接
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

// LogSwitch 日志开关
type LogSwitch struct {
	HttpAccess bool `json:"httpAccess"` // http 访问
	HttpError  bool `json:"httpError"`  // http 错误
	Mysql      bool `json:"Mysql"`      // Mysql 数据库
}

// Log 日志
type Log struct {
	Switch    LogSwitch    `json:"switch"`    // 启用
	Directory string       `json:"directory"` // 路径
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
	MaxAge         int64  `json:"maxAge"`     // 过期时间（秒）
	AutoRefreshAge int64  `json:"refreshAge"` // 自动重新刷新时间（秒）
}

// Config 配置
type Config struct {
	Mysql     Mysql     `json:"mysql"`     // Mysql
	Redis     Redis     `json:"redis"`     // Redis
	RateLimit RateLimit `json:"rateLimit"` // RateLimit
	Log       Log       `json:"log"`       // Log
	Request   Request   `json:"request"`   // Request
	Token     Token     `json:"token"`     // Token
}

// New 创建
func New(jsonFile string) (*Config, error) {
	type env struct {
		Dev  Config `json:"dev"`
		Prod Config `json:"prod"`
	}

	var e env
	if stream, err := os.ReadFile(jsonFile); err != nil {
		return nil, err
	} else if err = json.Unmarshal(stream, &e); err != nil {
		return nil, err
	}

	if IsDev() {
		return &e.Dev, nil
	} else {
		return &e.Prod, nil
	}
}
