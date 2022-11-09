/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-04-25 22:23:45
 */

package sms

// Type 验证码类型
type Type string

const (
	TypeLogin         Type = "Login"         // 登录
	TypeResetPassword Type = "ResetPassword" // 重置密码
)
