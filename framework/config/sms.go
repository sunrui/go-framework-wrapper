/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/09 14:18:09
 */

package config

// sms 配置对象
type sms struct {
	MagicCode     string `json:"magicCode"`     // 短信魔术码
	MaxAge        int    `json:"maxAge"`        // 过期时间（秒）
	MaxVerifyTime int    `json:"maxVerifyTime"` // 最多较验次数
	MaxSendPerDay int64  `json:"maxSendPerDay"` // 每日最多发送次数
}

// Sms 配置
func Sms() *sms {
	return &conf.Sms
}
