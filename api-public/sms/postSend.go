/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/03 17:51:03
 */

package sms

import (
	"encoding/json"
	"framework/app"
	"framework/config"
	"framework/proto/response"
	"framework/proto/result"
	"github.com/gin-gonic/gin"
	"service/core/sms"
)

// 发送验证码请求
type postCodeReq struct {
	Phone   string      `json:"phone" validate:"required,len=11,numeric"`                  // 手机号
	SmsType sms.SmsType `json:"smsType" validate:"required,oneof=Login," enums:"asc,desc"` // 验证码类型
}

// 发送验证码
func postSend(ctx *gin.Context) {
	var req postCodeReq

	// 较验参数
	app.ValidateParameter(ctx, &req)

	// 获取当天发送条数，判断是否超出最大条数限制
	count := sms.CountByPhoneAndDate(req.Phone, sms.GetNowDate())
	if count >= config.Sms().MaxSendPerDay {
		response.Result(ctx, result.RateLimit)
		return
	}

	// 创建 6 位验证码
	randomCode := sms.RandomCode()

	// 调用服务发送验证码
	channel, reqId, err := sms.Send(req.Phone, req.SmsType, randomCode)
	var comment string
	if err != nil {
		comment = err.Error()
	} else {
		// 备注对象
		type _comment struct {
			Chanel string `json:"chanel"`
			ReqId  string `json:"reqId"`
		}
		marshal, _ := json.Marshal(_comment{
			Chanel: channel,
			ReqId:  reqId,
		})

		comment = string(marshal)
	}

	// 存储发送记录
	history := sms.Sms{
		Phone:     req.Phone,
		SmsType:   req.SmsType,
		Code:      randomCode,
		Ip:        ctx.ClientIP(),
		UserAgent: ctx.Request.UserAgent(),
		Success:   err == nil,
		Comment:   comment,
	}
	history.Save()

	// 发送验证码失败
	if err != nil {
		response.Result(ctx, result.InternalError.WithData(err))
		return
	}

	// 将验证码缓存到 redis 中
	cache := sms.Cache{
		Phone:   req.Phone,
		SmsType: req.SmsType,
	}
	cache.SaveCode(randomCode)

	// 发送成功
	response.Result(ctx, result.Ok)
}
