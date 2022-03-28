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
	"medium-server-go/framework/result"
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
		app.Result(ctx, result.RateLimit)
		return
	}

	// 创建 6 位验证码
	sixNumber := sms.RandomCode()

	// 调用服务发送验证码
	channel, reqId, err := sms.Send(req.Phone, req.CodeType, sixNumber)
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
	code := sms.Code{
		Phone:     req.Phone,
		CodeType:  req.CodeType,
		Code:      sixNumber,
		Ip:        ctx.ClientIP(),
		UserAgent: ctx.Request.UserAgent(),
		Success:   err == nil,
		Comment:   comment,
	}
	code.Save()

	// 发送验证码失败
	if err != nil {
		app.Result(ctx, result.InternalError.WithData(err))
		return
	}

	// 将验证码缓存到 redis 中
	cache := sms.Cache{
		Phone:    req.Phone,
		CodeType: req.CodeType,
	}
	cache.Save(sms.CodeCache{
		Code:      sixNumber,
		ErrVerify: 0,
	})

	// 发送成功
	app.Result(ctx, result.Ok)
}

// 较验验证码
func postVerify(ctx *gin.Context) {
	var req postVerifyReq

	// 较验参数
	app.ValidateParameter(ctx, &req)

	// 缓存对象
	cache := sms.Cache{
		Phone:    req.Phone,
		CodeType: req.CodeType,
	}

	// 获取缓存数据
	if !cache.Exists() {
		app.Result(ctx, result.NotFound)
		return
	}

	// 较验验证码
	if !cache.Verify(req.Code) {
		app.Result(ctx, result.NotMatch)
		return
	}

	// 较验成功
	app.Result(ctx, result.Ok)
}
