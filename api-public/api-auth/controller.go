/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/01
 */
package api_auth

import (
	"github.com/gin-gonic/gin"
	"medium-server-go/common/app-gin"
	"medium-server-go/common/config"
	"medium-server-go/common/result"
	"medium-server-go/service/biz/sms"
	"medium-server-go/service/biz/user"
	"medium-server-go/service/provider"
	"net/http"
)

// 手机号码登录
func postLoginByPhone(ctx *gin.Context) {
	var req postLoginByPhoneReq

	// 较验参数
	errData, err := app_gin.ValidateParameter(ctx, &req)
	if err != nil {
		app_gin.Response(ctx, result.ParameterError.WithData(errData))
		return
	}

	// 如果非魔术验证码
	smsMagicCode := config.Get().Sms.MagicCode
	if smsMagicCode != "" && req.Code != smsMagicCode {
		// 短信缓存对象
		smsCache := sms.Cache{
			Phone:    req.Phone,
			CodeType: sms.CodeLogin,
		}

		// 获取缓存数据
		if !smsCache.Exists() {
			app_gin.Response(ctx, result.NotFound)
			return
		}

		// 较验验证码
		if !smsCache.Verify(req.Code) {
			app_gin.Response(ctx, result.NotMatch)
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
		user.SaveUser(userOne)
	}

	provider.Token.WriteToken(ctx, userOne.Id, 30*24*60*60)

	ctx.JSON(http.StatusOK,
		result.Ok.WithData(postLoginByPhoneRes{
			UserId: userOne.Id,
		}))
}

// 微信登录
func postLoginByWechat(ctx *gin.Context) {
	var req postLoginByPhoneReq

	// 较验参数
	errData, err := app_gin.ValidateParameter(ctx, &req)
	if err != nil {
		app_gin.Response(ctx, result.ParameterError.WithData(errData))
		return
	}
}

// 获取令牌
func getToken(ctx *gin.Context) {
	// 获取用户令牌
	tokenEntity, err := provider.Token.GetTokenEntity(ctx)
	if err != nil {
		app_gin.Response(ctx, result.NotFound)
		return
	}

	app_gin.Response(ctx, result.Ok.WithData(tokenEntity))
}

// 登出
func postLogout(ctx *gin.Context) {
	_, err := ctx.Cookie("token")
	if err != nil {
		app_gin.Response(ctx, result.NotFound)
		return
	}

	// 移除令牌
	provider.Token.RemoveToken(ctx)
	app_gin.Response(ctx, result.Ok)
}
