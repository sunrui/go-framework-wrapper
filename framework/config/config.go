/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/09 14:18:09
 */

package config

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"runtime"
)

// IsDebug 是否为调试环境
func IsDebug() bool {
	return true
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
var conf config

// 加载当前配置
func init() {
	var configFile = func() string {
		_, file, _, _ := runtime.Caller(0)
		path := filepath.Dir(file)

		var jsonFile string
		if IsDebug() {
			jsonFile = "config_debug.json"
		} else {
			jsonFile = "config_release.json"
		}

		return path + "/" + jsonFile
	}()

	if stream, err := ioutil.ReadFile(configFile); err != nil {
		panic(err.Error())
	} else if err = json.Unmarshal(stream, &conf); err != nil {
		panic(err.Error())
	}
}
