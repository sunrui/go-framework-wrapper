/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/03 17:51:03
 */

package sms

import (
	"framework/app"
	"framework/proto/result"
	"github.com/gin-gonic/gin"
	"service/sms"
)

// 较验验证码请求
type postVerifyReq struct {
	Phone   string   `json:"phone" validate:"required,len=11,numeric"` // 手机号
	SmsType sms.Type `json:"smsType" validate:"required,oneof=Login"`  // 验证码类型
	Code    string   `json:"code" validate:"required,len=6,numeric"`   // 验证码
}

// 较验验证码
func postVerify(ctx *gin.Context) result.Result {
	var req postVerifyReq

	// 较验参数
	app.ValidateParameter(ctx, &req)

	// 缓存对象
	cache := sms.Cache{
		Phone:   req.Phone,
		SmsType: req.SmsType,
	}

	// 获取缓存数据
	if !cache.Exists() {
		return result.NotFound
	}

	// 较验验证码
	if !cache.Verify(req.Code) {
		return result.NotFound
	}

	// 较验成功
	return result.Ok
}