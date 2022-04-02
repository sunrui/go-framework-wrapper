package auth

import (
	"github.com/gin-gonic/gin"
	"medium-server-go/enum"
	"medium-server-go/framework/app"
	"medium-server-go/framework/config"
	"medium-server-go/framework/proto/response"
	"medium-server-go/framework/proto/result"
	"medium-server-go/framework/proto/token"
	"medium-server-go/service/sms"
	"medium-server-go/service/user"
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
	smsMagicCode := config.Get().Sms.MagicCode
	if smsMagicCode != "" && req.Code != smsMagicCode {
		// 短信缓存对象
		smsCache := sms.Cache{
			Phone:   req.Phone,
			SmsType: enum.SmsLogin,
		}

		// 获取缓存数据
		if !smsCache.Exists() {
			response.Response(ctx).Data(result.NotFound.WithData(req))
			return
		}

		// 较验验证码
		if !smsCache.Verify(req.Code) {
			response.Response(ctx).Data("result.NotMatch")
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

	response.Response(ctx).Data(postLoginByPhoneRes{
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
		response.Response(ctx).Data(result.NotFound)
		return
	}

	response.Response(ctx).Data(tokenEntity)
}

// 登出
func postLogout(ctx *gin.Context) {
	_, err := ctx.Cookie("token")
	if err != nil {
		response.Response(ctx).Data(result.NotFound.WithData(err.Error()))
		return
	}

	// 移除令牌
	token.Remove(ctx)
	response.Response(ctx).Ok()
}
