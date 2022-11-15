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

type Log struct {
	Level         string `json:"level"`         // 等级
	EnableFile    bool   `json:"enableFile"`    // 写入文件
	EnableConsole bool   `json:"enableConsole"` // 写入控制台
}
