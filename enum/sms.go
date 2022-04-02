/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022/04/03 00:12:03
 */

package enum

// SmsType 验证码类型
type SmsType string

const (
	SmsLogin         = "SmsLogin"       // 登录
	SmsResetPassword = "RESET_PASSWORD" // 重置密码
)
