/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/20 21:25:20
 */

package enum

// CodeType 验证码类型
type CodeType string

const (
	CodeLogin         = "LOGIN"          // 登录
	CodeResetPassword = "RESET_PASSWORD" // 重置密码
)
