/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/09 14:18:09
 */

package config

import (
	"encoding/json"
	"flag"
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

// 配置对象
type config struct {
	Log       log       `json:"log"`       // Log 配置对象
	RateLimit rateLimit `json:"rateLimit"` // RateLimit 配置对象
	Swagger   swagger   `json:"swagger"`   // Swagger 配置对象
	Mysql     mysql     `json:"mysql"`     // Mysql 配置对象
	Redis     redis     `json:"redis"`     // Redis 配置对象
	Jwt       jwt       `json:"jwt"`       // Jwt 配置对象
	Sms       sms       `json:"sms"`       // Sms 配置对象
}

// 当前配置
var conf *config

// Get 获取当前配置
func Get() *config {
	if conf != nil {
		return conf
	}

	// 配置文件
	var configFile = func() string {
		_, file, _, _ := runtime.Caller(0)
		path := filepath.Dir(file)

		var jsonFile string
		if IsDev() {
			jsonFile = "config_dev.json"
		} else {
			jsonFile = "config_product.json"
		}

		return path + "/" + jsonFile
	}()

	// 加载配置文件
	conf = &config{}
	if stream, err := os.ReadFile(configFile); err != nil {
		panic(err.Error())
	} else if err = json.Unmarshal(stream, conf); err != nil {
		panic(err.Error())
	}

	return conf
}

// 当前环境
var build *string

// IsDev 是否为开发环境
func IsDev() bool {
	return build != nil && *build != "product"
}

// IsProduct 是否为生产环境
func IsProduct() bool {
	return !IsDev()
}

// 初始化
func init() {
	// 解析参数，如 -build product
	flag.Parse()
	build = flag.String("build", "dev", "编译类型")
}
