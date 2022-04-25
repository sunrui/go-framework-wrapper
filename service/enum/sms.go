/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-04-25 22:23:45
 */

package enum

// SmsType 验证码类型
type SmsType string

const (
	Login         SmsType = "Login"         // 登录
	ResetPassword SmsType = "ResetPassword" // 重置密码
)
