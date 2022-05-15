/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022/04/03 13:44:03
 */

package common

import (
	"framework/proto/response"
	"framework/proto/result"
	"github.com/gin-gonic/gin"
	"time"
)

var build = time.Now()

// @Summary  编译时间
// @Tags     通用
// @Produce  json
// @Success  200  {object}  result.Result{data=string}  true  {"status":"Ok","description":"成功"}
// @Router   /public/common/build [get]
func getBuild(ctx *gin.Context) {
	response.New(ctx).Data(build.Format("2006-01-02 15:04:05"))
}

// @Summary  结果列表
// @Tags     通用
// @Produce  json
// @Success  200  {object}  result.Result  true  {"status":"Ok","description":"成功"}
// @Success  400  {object}  result.Result  true  {"status":"BadRequest","description":"语法错误"}
// @Success  401  {object}  result.Result  true  {"status":"NoAuth","description":"没有登录"}
// @Success  403  {object}  result.Result  true  {"status":"Forbidden","description":"没有权限"}
// @Success  404  {object}  result.Result  true  {"status":"NotFound","description":"不存在"}
// @Success  405  {object}  result.Result  true  {"status":"MethodNotAllowed","description":"请求方式不允许"}
// @Success  409  {object}  result.Result  true  {"status":"Conflict","description":"请求冲突"}
// @Success  429  {object}  result.Result  true  {"status":"RateLimit","description":"限流"}
// @Success  500  {object}  result.Result  true  {"status":"InternalError","description":"内部错误"}
// @Success  501  {object}  result.Result  true  {"status":"NotImplemented","description":"未实现"}
// @Success  502  {object}  result.Result  true  {"status":"BadGateway","description":"网关错误"}
// @Router   /public/common/result [get]
func getResult(ctx *gin.Context) {
	response.New(ctx).Data(result.All())
}
