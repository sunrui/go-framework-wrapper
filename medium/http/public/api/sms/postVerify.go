/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-29 22:24:39
 */

package sms

import (
	"framework/app"
	"framework/result"
	"github.com/gin-gonic/gin"
	"medium/service/user"
)

// 手机号码登录请求
type postVerifyReq struct {
	Phone       string          `json:"phone" validate:"required,len=11"`       // 手机号
	Code        string          `json:"code" validate:"required,len=6,numeric"` // 验证码
	DeviceType  user.DeviceType `json:"deviceType" validate:"required"`         // 设备类型
	PackageName string          `json:"packageName" validate:"required"`        // 包名
	AppVersion  string          `json:"appVersion" validate:"required"`         // 软件版本
}

// 手机号码登录结果
type postVerifyRes struct {
	UserId string `json:"userId"` // 用户 id
}

// @Summary         登录 - 手机
// @Tags            认证
// @Accept          json
// @Produce         json
// @Param           "req"  body  postVerifyReq  true  "req"
// @ApprovalSuccess 200    {object}  postVerifyRes
// @Failure         400 {object} result.Result
// @Router          /auth/sms/verify [post]
func postVerify(ctx *gin.Context) *result.Result {
	var r postVerifyReq

	// 较验参数
	server.ValidateParameter(ctx, &r)

	// 返回结果
	return nil
}
