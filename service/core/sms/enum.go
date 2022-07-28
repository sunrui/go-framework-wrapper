/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-04-25 22:23:45
 */

package sms

// SmsType 验证码类型
type SmsType string

const (
	SmsTypeLogin         SmsType = "SmsTypeLogin"         // 登录
	SmsTypeResetPassword SmsType = "SmsTypeResetPassword" // 重置密码
)
