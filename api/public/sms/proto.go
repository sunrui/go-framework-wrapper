/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/03 17:51:03
 */

package sms

import (
	"medium-server-go/enum"
)

// 发送验证码请求
type postCodeReq struct {
	Phone   string       `json:"phone" validate:"required,len=11,numeric"`                 // 手机号
	SmsType enum.SmsType `json:"smsType" validate:"required,oneof=LOGIN" enums:"asc,desc"` // 验证码类型
}

// 较验验证码请求
type postVerifyReq struct {
	Phone   string       `json:"phone" validate:"required,len=11,numeric"` // 手机号
	SmsType enum.SmsType `json:"smsType" validate:"required,oneof=LOGIN"`  // 验证码类型
	Code    string       `json:"code" validate:"required,len=6,numeric"`   // 验证码
}
