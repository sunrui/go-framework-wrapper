/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-07 00:15:14
 */

package config

import "github.com/sirupsen/logrus"

// 数据库
type mysql struct {
	User         string `json:"user"`         // 用户名
	Password     string `json:"password"`     // 密码
	Host         string `json:"host"`         // 主机
	Port         int    `json:"port"`         // 端口
	Database     string `json:"database"`     // 数据库
	MaxOpenConns int    `json:"maxOpenConns"` // 最大打开连接
	MaxIdleConns int    `json:"maxIdleConns"` // 最大空闲连接
}

// 限流
type rateLimit struct {
	Capacity int64 `json:"capacity"` // 令牌桶容量
	Quantum  int64 `json:"quantum"`  // 每隔多少秒
}

// 缓存
type redis struct {
	Host     string `json:"host"`     // 主机
	Port     int    `json:"port"`     // 端口
	Password string `json:"password"` // 密码
	Database int    `json:"database"` // 数据库
}

type logSwitch struct {
	HttpAccess bool `json:"httpAccess"` // http 访问
	HttpError  bool `json:"httpError"`  // http 错误
	Mysql      bool `json:"mysql"`      // mysql
}

// 日志
type log struct {
	Switch    logSwitch    `json:"switch"`    // 启用
	Directory string       `json:"directory"` // 路径
	Level     logrus.Level `json:"level"`     // 等级
}

// 请求
type request struct {
	Dump bool `json:"dump"` // 导出
}

// 令牌
type token struct {
	JwtSecret      string `json:"jwtSecret"`  // jwt 密钥
	Key            string `json:"key"`        // 键名
	MaxAge         int64  `json:"max_age"`    // 过期时间（秒）
	AutoRefreshAge int64  `json:"refreshAge"` // 自动重新刷新时间（秒）
}
