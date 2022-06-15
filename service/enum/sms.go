/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-04-25 22:23:45
 */

package enum

// SmsType 验证码类型
type SmsType string

const (
	SmsLogin         SmsType = "SmsLogin"         // 登录
	SmsResetPassword SmsType = "SmsResetPassword" // 重置密码
)
