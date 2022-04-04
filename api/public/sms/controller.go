/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/03 18:07:03
 */

package sms

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"medium-server-go/framework/app"
	"medium-server-go/framework/proto/response"
	"medium-server-go/framework/proto/result"
	"medium-server-go/service/sms"
)

// 发送验证码
func postCode(ctx *gin.Context) {
	var req postCodeReq

	// 较验参数
	app.ValidateParameter(ctx, &req)

	// 获取当天发送条数，判断是否超出最大条数限制
	count := sms.CountByPhoneAndDate(req.Phone, sms.GetNowDate())
	if count >= 5 {
		response.New(ctx).Data(result.RateLimit)
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
		response.New(ctx).Data(result.InternalError.WithData(err))
		return
	}

	// 将验证码缓存到 redis 中
	cache := sms.Cache{
		Phone:   req.Phone,
		SmsType: req.SmsType,
	}
	cache.SaveCode(randomCode)

	// 发送成功
	response.New(ctx).Ok()
}

// 较验验证码
func postVerify(ctx *gin.Context) {
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
		response.New(ctx).Data(result.NotFound)
		return
	}

	// 较验验证码
	if !cache.Verify(req.Code) {
		response.New(ctx).Data(result.NotMatch)
		return
	}

	// 较验成功
	response.New(ctx).Ok()
}
