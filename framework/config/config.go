/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/09 14:18:09
 */

package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
)

// log 配置对象
type log struct {
	Enable bool   `json:"enable"` // 开关
	Level  string `json:"level"`  // 等级
}

// rateLimit 配置对象
type rateLimit struct {
	Quantum  int64 `json:"quantum"`  // 间隔时间（秒）
	Capacity int64 `json:"capacity"` // 令牌桶容量
}

// swagger 配置对象
type swagger struct {
	Enable bool `json:"enable"` // 是否启用
}

// mysql 配置对象
type mysql struct {
	Host     string `json:"host"`     // 主机
	Port     int    `json:"port"`     // 端口
	Database string `json:"database"` // 数据库
	User     string `json:"user"`     // 用户名
	Password string `json:"password"` // 密码
}

// redis 配置对象
type redis struct {
	Host     string `json:"host"`     // 主机
	Port     int    `json:"port"`     // 端口
	Password string `json:"password"` // 密码
	Database int    `json:"database"` // 数据库
	Timeout  int    `json:"timeout"`  // 超时时间（秒）
}

// jwt 配置对象
type jwt struct {
	Key         string `json:"key"`         // 主键
	Secret      []byte `json:"secret"`      // 密钥
	MaxAge      int    `json:"maxAge"`      // 过期时间（秒）
	AutoRefresh int    `json:"autoRefresh"` // 自动续订（秒）
}

// sms 配置对象
type sms struct {
	MagicCode     string `json:"magicCode"`     // 短信魔术码
	MaxAge        int    `json:"maxAge"`        // 过期时间（秒）
	MaxVerifyTime int    `json:"maxVerifyTime"` // 最多较验次数
	MaxSendPerDay int64  `json:"maxSendPerDay"` // 每日最多发送次数
}

// Config 配置对象
type Config struct {
	Log       log       `json:"log"`       // Log 配置对象
	RateLimit rateLimit `json:"rateLimit"` // RateLimit 配置对象
	Swagger   swagger   `json:"swagger"`   // Swagger 配置对象
	Mysql     mysql     `json:"mysql"`     // Mysql 配置对象
	Redis     redis     `json:"redis"`     // Redis 配置对象
	Jwt       jwt       `json:"jwt"`       // Jwt 配置对象
	Sms       sms       `json:"sms"`       // Sms 配置对象
}

// 当前配置
var cur *Config

// Get 获取当前配置
func Get() *Config {
	if cur != nil {
		return cur
	}

	// 配置文件
	var configFile = func() string {
		_, file, _, _ := runtime.Caller(0)
		path := filepath.Dir(file)

		var jsonFile string
		if IsDev() {
			jsonFile = "config_dev.json"
		} else {
			jsonFile = "config_prod.json"
		}

		return path + "/" + jsonFile
	}()

	// 加载配置文件
	cur = &Config{}
	if stream, err := os.ReadFile(configFile); err != nil {
		panic(err.Error())
	} else if err = json.Unmarshal(stream, cur); err != nil {
		panic(err.Error())
	}

	return cur
}
