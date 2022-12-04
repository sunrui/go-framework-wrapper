/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-29 22:24:39
 */

package common

import (
	"framework/app/result"
	"framework/server"
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

// @Summary 编译时间3
// @Tags    通用3
// @Produce json
// @Success 200   {object} result.Result
// @Param   email body     string true "message/rfc822" SchemaExample(Subject: Testmail\r\n\r\nBody Message\r\n)
// @Router  /public/common/send [get]
func postSend(ctx *gin.Context) *result.Result {
	var r postSendReq

	// 较验参数
	server.ValidateParameter(ctx, &r)

	// 返回结果
	return nil
}
