/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/09 14:18:09
 */

package config

// jwt 配置对象
type jwt struct {
	Key         string `json:"key"`         // 主键
	Secret      []byte `json:"secret"`      // 密钥
	MaxAge      int    `json:"maxAge"`      // 过期时间（秒）
	AutoRefresh int    `json:"autoRefresh"` // 自动续订（秒）
}

// Jwt 配置
func Jwt() *jwt {
	return &conf.Jwt
}
