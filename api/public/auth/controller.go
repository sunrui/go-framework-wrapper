package auth

import (
	"github.com/gin-gonic/gin"
	"medium-server-go/enum"
	"medium-server-go/framework/app"
	"medium-server-go/framework/config"
	"medium-server-go/framework/result"
	"medium-server-go/framework/token"
	"medium-server-go/service/sms"
	"medium-server-go/service/user"
)

// 手机号码登录
func postLoginByPhone(ctx *gin.Context) {
	var req postLoginByPhoneReq

	// 较验参数
	app.ValidateParameter(ctx, &req)

	// 如果非魔术验证码
	smsMagicCode := config.Get().Sms.MagicCode
	if smsMagicCode != "" && req.Code != smsMagicCode {
		// 短信缓存对象
		smsCache := sms.Cache{
			Phone:    req.Phone,
			CodeType: enum.CodeLogin,
		}

		// 获取缓存数据
		if !smsCache.Exists() {
			app.Response(ctx, result.NotFound.WithKeyPair("code", req.Code))
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

	token.Write(ctx, userOne.Id, 30*24*60*60)

	app.Response(ctx, result.Ok.WithData(postLoginByPhoneRes{
		UserId: userOne.Id,
	}))
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
	token.Remove(ctx)
	app.Response(ctx, result.Ok)
}
