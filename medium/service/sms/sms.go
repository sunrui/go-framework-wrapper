/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-29 22:09:16
 */

package sms

import "framework/app/mysql"

// Type 短信类型
type Type string

const (
	LoginType Type = "LOGIN" // 登录类型
)

// Sms 短信
type Sms struct {
	mysql.Model        // 通用参数
	Phone       string `json:"phone" gorm:"comment:手机号"`        // 手机号
	Type        Type   `json:"type" gorm:"comment:短信类型"`        // 短信类型
	Code        string `json:"code" gorm:"comment:验证码"`         // 验证码
	Ip          string `json:"ip" gorm:"comment:ip 地址"`         // ip 地址
	UserAgent   string `json:"userAgent" gorm:"comment:用户 ua"`  // 用户 ua
	IsSuccess   bool   `json:"IsSuccess" gorm:"comment:是否发送成功"` // 是否成功
	Comment     string `json:"comment" gorm:"comment:备注"`       // 备注
}
