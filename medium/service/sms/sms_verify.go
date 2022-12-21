/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-29 22:09:16
 */

package sms

import "framework/app/mysql"

// Verify 短信较验
type Verify struct {
	mysql.Model        // 通用参数
	Phone       string `json:"phone" gorm:"comment:手机号"`       // 手机号
	Type        Type   `json:"type" gorm:"comment:短信类型"`       // 短信类型
	Code        string `json:"code" gorm:"comment:验证码"`        // 验证码
	InputCode   string `json:"inputCode" gorm:"comment:输入验证码"` // 输入验证码
}

func (Verify) TableName() string {
	return "t_sms_verify"
}
