/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/03 17:51:03
 */

package sms

import (
	"framework/app"
	"framework/config"
	"framework/db"
	"framework/proto/result"
	"github.com/gin-gonic/gin"
	"service/core/sms"
)

// 发送验证码请求
type postCodeReq struct {
	Phone   string   `json:"phone" validate:"required,len=11,numeric"`              // 手机号
	SmsType sms.Type `json:"smsType" validate:"required,oneof=Login" enums:"Login"` // 验证码类型
}

// 发送验证码
func postSend(ctx *gin.Context) result.Result {
	var req postCodeReq

	// 较验参数
	app.ValidateParameter(ctx, &req)

	// 获取当天发送条数，判断是否超出最大条数限制
	count := sms.CountByPhoneAndDate(req.Phone, sms.GetNowDate())
	if count >= config.Sms().MaxSendPerDay {
		return result.RateLimit
	}

	// 创建 6 位验证码
	randomCode := sms.RandomCode()

	// 调用服务发送验证码
	err := sms.Send(req.Phone, req.SmsType, randomCode)

	// 保存发送记录
	smsOne := sms.Sms{
		Phone:     req.Phone,
		Type:      req.SmsType,
		Code:      randomCode,
		Ip:        ctx.ClientIP(),
		UserAgent: ctx.Request.UserAgent(),
		Success:   err == nil,
		Comment: func() string {
			if err != nil {
				return err.Error()
			} else {
				return ""
			}
		}(),
	}

	if tx := db.Mysql.Save(&smsOne); tx.Error != nil {
		panic(tx.Error.Error())
	}

	// 发送验证码失败
	if err != nil {
		return result.InternalError.WithData(err)
	}

	// 将验证码缓存到 redis 中
	cache := sms.Cache{
		Phone:   req.Phone,
		SmsType: req.SmsType,
	}
	cache.SaveCode(randomCode)

	// 发送成功
	return result.Ok
}
