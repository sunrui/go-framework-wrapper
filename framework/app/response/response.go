/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-24 22:35:55
 */

package response

import (
	"fmt"
	"framework/app/log"
	"framework/app/request"
	"framework/app/token"
	"framework/config"
	"framework/result"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 获取结果数据
func getResultBuffer(ctx *gin.Context, r *result.Result) string {
	// 未启用日志返回
	if !config.Inst().Log.Enable {
		return ""
	}

	// 获取 request 对象
	req := request.Get(ctx)

	// method http://host:port?query protocol
	buffer := req.Method + " " + req.Uri + " " + req.Proto

	// userId
	if userId := token.GetUserId(ctx); userId != nil {
		buffer += " - userId(" + *userId + ")"
	}

	// 空一行
	buffer += "\n"

	// header
	for _, header := range req.Headers {
		buffer += header + "\n"
	}

	// 空一行
	buffer += "\n"

	// body
	if req.Body != nil {
		buffer += *req.Body + "\n"
	} else {
		buffer += "<null>\n"
	}

	// 空一行
	buffer += "\n"

	// 结果
	buffer += r.String() + "\n"

	// 空一行
	buffer += "\n"

	return buffer
}

// Response 返回
func Response(ctx *gin.Context, r *result.Result) {
	// 结果导出请求
	if request.IsDump(ctx) {
		req := request.Get(ctx)
		r.Request = &req
	}

	go func() {
		// 开启控制台
		if config.IsDev() {
			fmt.Println(r)
		}

		// 记录日志
		buffer := getResultBuffer(ctx, r)
		if r.Code == result.Ok.Code {
			log.Http.Debugln(buffer)
		} else {
			log.Http.Errorln(buffer)
		}
	}()

	// 返回客户端
	ctx.AbortWithStatusJSON(http.StatusOK, r)
}
