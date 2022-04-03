package example

import (
	"github.com/gin-gonic/gin"
	"medium-server-go/framework/app"
	"medium-server-go/framework/proto/request"
	"medium-server-go/framework/proto/response"
	"medium-server-go/framework/proto/result"
)

// @Summary   发送验证码
// @Tags     演示
// @Accept   json
// @Produce  json
// @Param     json  body      postSmsReq     true  "struct"
// @Success   200   {object}  result.Result  true  {"code":"Ok","message":"成功"}
// @Response  201   {object}  result.Result  true  {"code":"NotMatch","message":"不匹配"}
// @Router    /sms [post]
func postSms(ctx *gin.Context) {
	var req postSmsReq

	// 较验参数
	app.ValidateParameter(ctx, &req)

	// 检测是否已经发送
	smsCode := find(req.Phone)
	if smsCode != nil {
		response.Response(ctx).Data(result.Duplicate)
		return
	}

	// 发送验证码
	create(req.Phone)

	r := result.Ok.WithMessage("发送成功了")
	println(r.String())

	// 假定发送成功
	response.Response(ctx).Data(r)
}

// @Summary  获取所有验证码
// @Tags      演示
// @Accept    json
// @Produce   json
// @Param    page      query     int                            false  "分页"
// @Param    pageSize  query     int                            false  "分页大小"
// @Success  200       {object}  result.PageResult{data=[]Sms}  true
// @Router   /sms [get]
func getSms(ctx *gin.Context) {
	var req request.PageRequest

	// 较验参数
	app.ValidateParameter(ctx, &req)

	response.Response(ctx).PageData(all(), result.Pagination{
		Page:      req.Page,
		PageSize:  req.PageSize,
		TotalPage: 10,
		TotalSize: 100,
	})
}

// @Summary  获取手机号获取验证码
// @Description
// @Tags     演示
// @Accept   json
// @Produce  json
// @Param    phone        path      string                   true  "13012341234"
// @Success  200          {object}  result.Result{data=Sms}  true
// @Failure  300          {object}  result.Result            true  "{"code":"NotFound", "message":"不存在"}"
// @Failure  301          {object}  result.Result            true  "{"code":"NotMatch", "message":"不匹配"}"
// @Header   400,401,402  {string}  Token2                   "token2"
// @Router   /sms/{phone} [get]
func getSmsOne(ctx *gin.Context) {
	phone := ctx.Param("phone")

	// 检测是否已经发送
	smsCode := find(phone)
	if smsCode == nil {
		response.Response(ctx).Data(result.NotFound)
		return
	}

	// 返回验证码
	response.Response(ctx).Data(smsCode)
}
