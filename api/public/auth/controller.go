/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/01
 */
package auth

import (
	"github.com/gin-gonic/gin"
	"medium-server-go/framework/app"
	"medium-server-go/framework/config"
	"medium-server-go/framework/result"
	"medium-server-go/framework/token"
	sms2 "medium-server-go/service/sms"
	user2 "medium-server-go/service/user"
	"net/http"
)

// 手机号码登录
func postLoginByPhone(ctx *gin.Context) {
	var req postLoginByPhoneReq

	// 较验参数
	errData, err := app.ValidateParameter(ctx, &req)
	if err != nil {
		app.Response(ctx, result.ParameterError.WithData(errData))
		return
	}

	// 如果非魔术验证码
	smsMagicCode := config.Get().Sms.MagicCode
	if smsMagicCode != "" && req.Code != smsMagicCode {
		// 短信缓存对象
		smsCache := sms2.Cache{
			Phone:    req.Phone,
			CodeType: sms2.CodeLogin,
		}

		// 获取缓存数据
		if !smsCache.Exists() {
			app.Response(ctx, result.NotFound)
			return
		}

		// 较验验证码
		if !smsCache.Verify(req.Code) {
			app.Response(ctx, result.NotMatch)
			return
		}

		// 移除验证码
		smsCache.Del()
	}

	// 查找当前用户是否存在
	userOne := user2.FindByPhone(req.Phone)
	if userOne == nil {
		userOne = &user2.User{
			Phone:     req.Phone,
			Ip:        ctx.ClientIP(),
			UserAgent: ctx.Request.UserAgent(),
		}

		// 创建新的用户
		user2.SaveUser(userOne)
	}

	token.WriteToken(ctx, userOne.Id, 30*24*60*60)

	ctx.JSON(http.StatusOK,
		result.Ok.WithData(postLoginByPhoneRes{
			UserId: userOne.Id,
		}))
}

// 微信登录
func postLoginByWechat(ctx *gin.Context) {
	var req postLoginByPhoneReq

	// 较验参数
	errData, err := app.ValidateParameter(ctx, &req)
	if err != nil {
		app.Response(ctx, result.ParameterError.WithData(errData))
		return
	}
}

// 获取令牌
func getToken(ctx *gin.Context) {
	// 获取用户令牌
	tokenEntity, err := token.GetTokenEntity(ctx)
	if err != nil {
		app.Response(ctx, result.NotFound)
		return
	}

	app.Response(ctx, result.Ok.WithData(tokenEntity))
}

// 登出
func postLogout(ctx *gin.Context) {
	_, err := ctx.Cookie("token")
	if err != nil {
		app.Response(ctx, result.NotFound)
		return
	}

	// 移除令牌
	token.RemoveToken(ctx)
	app.Response(ctx, result.Ok)
}
