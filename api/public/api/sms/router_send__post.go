/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-29 22:24:39
 */

package sms

import (
	"framework/app/result"
	"framework/app/server"
	"github.com/gin-gonic/gin"
	"medium/service/user"
)

// 手机号码登录请求
type postSendReq struct {
	Phone       string          `json:"phone" validate:"required,len=11"`       // 手机号
	Code        string          `json:"code" validate:"required,len=6,numeric"` // 验证码
	DeviceType  user.DeviceType `json:"deviceType" validate:"required"`         // 设备类型
	PackageName string          `json:"packageName" validate:"required"`        // 包名
	AppVersion  string          `json:"appVersion" validate:"required"`         // 软件版本
}

// 手机号码登录结果
type postSendRes struct {
	UserId string `json:"userId"` // 用户 id
}

// @Summary         短信 - 发送
// @Tags            认证
// @Accept          json
// @Produce         json
// @Param           "req"  body  postSendReq  true  " "
// @ApprovalSuccess 200    {object}  postSendRes
// @Failure         400 {object} result.Result
// @Router          /auth/sms/send [post]
func postSend(ctx *gin.Context) result.Result {
	var r postSendReq

	// 较验参数
	server.ValidateParameter(ctx, &r)

	// 返回结果
	return result.Ok
}
