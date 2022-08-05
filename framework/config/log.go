/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/09 14:18:09
 */

package config

// log 配置相关
type log struct {
	Enable bool   `json:"enable"` // 开关
	Level  string `json:"level"`  // 等级
}

// Log 配置
func Log() *log {
	return &conf.Log
}
