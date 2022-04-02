package demo

import (
	"github.com/gin-gonic/gin"
	"medium-server-go/framework/app"
	"medium-server-go/framework/result"
)

// @Summary   发送验证码
// @Tags      演示
// @Accept    json
// @Produce   json
// @Param     "req"  body      postSmsReq     true  "req"
// @Success   200    {object}  result.Result  true  {"code":"response","message":"成功"}
// @Response  201    {object}  result.Result  true  {"code":"Response","message":"成功"}
// @Response  202    {object}  result.Result  true  {"code":"Response2","message":"成功"}
// @Response  203    {object}  result.Result  true  {"code":"Response3","message":"成功"}
// @Router    /sms [post]
func postSms(ctx *gin.Context) {
	var req postSmsReq

	// 较验参数
	app.ValidateParameter(ctx, &req)

	// 检测是否已经发送
	smsCode := find(req.Phone)
	if smsCode != nil {
		app.Response(ctx).Data(result.Duplicate)
		return
	}

	// 发送验证码
	create(req.Phone)

	// 假定发送成功
	app.Response(ctx).Ok()
}

// @Summary  获取所有验证码
// @Tags     演示
// @Accept   json
// @Produce  json
// @Param    page      query     int    false  "分页"
// @Param    pageSize  query     int    false  "分页大小"
// @Success  200       {object}  result.PageResult  true
// @Router   /sms [get]
func getSms(ctx *gin.Context) {
	var req app.PageRequest

	// 较验参数
	app.ValidateParameter(ctx, &req)

	app.Response(ctx).Data(all())
}

// @Summary  获取某个验证码
// @Tags     演示
// @Accept   json
// @Produce  json
// @Param    phone  path      string  true  "13012341234"
// @Success  200    {object}  Sms     true
// @Router   /sms/{phone} [get]
func getSmsOne(ctx *gin.Context) {
	phone := ctx.Param("phone")

	// 检测是否已经发送
	smsCode := find(phone)
	if smsCode == nil {
		app.Response(ctx).Data(result.NotFound)
		return
	}

	// 返回验证码
	app.Response(ctx).Data(smsCode)
}
