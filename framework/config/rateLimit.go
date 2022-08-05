/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/09 14:18:09
 */

package config

// rateLimit 配置对象
type rateLimit struct {
	Quantum  int64 `json:"quantum"`  // 间隔时间（秒）
	Capacity int64 `json:"capacity"` // 令牌桶容量
}

// RateLimit 配置
func RateLimit() *rateLimit {
	return &conf.RateLimit
}
