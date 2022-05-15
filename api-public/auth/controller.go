/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-04-25 22:14:59
 */

package auth

import (
	"framework/app"
	"framework/config"
	"framework/proto/response"
	"framework/proto/result"
	"framework/proto/token"
	"github.com/gin-gonic/gin"
	"service/core/sms"
	"service/core/user"
	"service/enum"
)

// @Summary  登录 - 手机
// @Tags     认证
// @Accept   json
// @Produce  json
// @Param    "req"  body      postLoginByPhoneReq  true  "req"
// @Success  200    {object}  postLoginByPhoneRes
// @Failure  400    {object}  result.Result
// @Router   /auth/login/phone [post]
func postLoginByPhone(ctx *gin.Context) {
	var req postLoginByPhoneReq

	// 较验参数
	app.ValidateParameter(ctx, &req)

	// 如果非魔术验证码
	smsMagicCode := config.Sms().MagicCode
	if smsMagicCode != "" && req.Code != smsMagicCode {
		// 短信缓存对象
		smsCache := sms.Cache{
			Phone:   req.Phone,
			SmsType: enum.Login,
		}

		// 获取缓存数据
		if !smsCache.Exists() {
			response.New(ctx).Data(result.NotFound.WithData(req))
			return
		}

		// 较验验证码
		if !smsCache.Verify(req.Code) {
			response.New(ctx).Data("common.NotMatch")
			return
		}

		// 移除验证码
		smsCache.Del()
	}

	// 查找当前用户是否存在
	userOne := user.FindByPhone(req.Phone)
	if userOne == nil {
		userOne = &user.User{
			Phone:     req.Phone,
			Ip:        ctx.ClientIP(),
			UserAgent: ctx.Request.UserAgent(),
		}

		// 创建新的用户
		userOne.Save()
	}

	token.Write(ctx, userOne.Id)

	response.New(ctx).Data(postLoginByPhoneRes{
		UserId: userOne.Id,
	})
}

// 微信登录
func postLoginByWechat(ctx *gin.Context) {
	var req postLoginByPhoneReq

	// 较验参数
	app.ValidateParameter(ctx, &req)
}

// 获取令牌
func getToken(ctx *gin.Context) {
	// 获取用户令牌
	tokenEntity, err := token.Get(ctx)
	if err != nil {
		response.New(ctx).Data(result.NotFound)
		return
	}

	response.New(ctx).Data(tokenEntity)
}

// 登出
func postLogout(ctx *gin.Context) {
	_, err := ctx.Cookie("token")
	if err != nil {
		response.New(ctx).Data(result.NotFound.WithData(err.Error()))
		return
	}

	// 移除令牌
	token.Remove(ctx)
	response.New(ctx).Ok()
}
