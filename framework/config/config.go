/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/09 14:18:09
 */

package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
)

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
}

// jwt 配置对象
type jwt struct {
	Secret []byte `json:"secret"` // 密钥
}

// sms 配置对象
type sms struct {
	MagicCode string `json:"magicCode"` // 短信魔术码
}

// Config 配置对象
type Config struct {
	Mysql mysql `json:"mysql"` // Mysql 配置对象
	Redis redis `json:"redis"` // Redis 配置对象
	Jwt   jwt   `json:"jwt"`   // jwt 配置对象
	Sms   sms   `json:"sms"`   // sms 配置对象
}

// json 反射对象
type jsonConfig struct {
	Environment string `json:"environment"` // 当前环境
	Debug       Config `json:"debug"`       // 开发环境
	Release     Config `json:"release"`     // 正式环境
}

// 当前配置
var config *jsonConfig

// IsDebugMode 是否在调试环境
func IsDebugMode() bool {
	return strings.ToLower(config.Environment) == "debug"
}

// Get 获取当前配置
func Get() *Config {
	if IsDebugMode() {
		return &config.Debug
	} else {
		return &config.Release
	}
}

// 加载当前配置
func init() {
	var stream []byte
	var err error

	// 获取当前项目根目录 config.json
	pwd, _ := os.Getwd()
	if stream, err = ioutil.ReadFile(pwd + "/config.json"); err != nil {
		panic(err.Error())
	}

	// 反射配置文件
	if err = json.Unmarshal(stream, &config); err != nil {
		panic(err.Error())
	}
}
