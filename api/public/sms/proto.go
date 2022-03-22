/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/03 17:51:03
 */

package sms

import (
	"medium-server-go/service/sms"
)

// 发送验证码请求
type postCodeReq struct {
	Phone    string       `json:"phone" binding:"required" validate:"len=11,numeric"` // 手机号
	CodeType sms.CodeType `json:"codeType" binding:"required" validate:"oneof=LOGIN"` // 验证码类型
}

// 较验验证码请求
type postVerifyReq struct {
	Phone    string       `json:"phone" binding:"required" validate:"len=11,numeric"` // 手机号
	CodeType sms.CodeType `json:"codeType" binding:"required" validate:"oneof=LOGIN"` // 验证码类型
	Code     string       `json:"code" binding:"required" validate:"len=6,numeric"`   // 验证码
}
