/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-12-04 09:25:52
 */

package server

import (
	"fmt"
	"framework/app/glog"
	"framework/app/result"
	"framework/server/middleware"
	request2 "framework/server/request"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 获取结果数据
func (server Server) getResultBuffer(ctx *gin.Context, r *result.Result, elapsed int64) string {
	// 获取 request
	req := request2.Get(ctx)

	var buffer = fmt.Sprintf("%dms - userId(%s)", elapsed, func() string {
		if payload, _, err := server.token.GetPayload(ctx); err == nil {
			return payload.UserId
		} else {
			return "nil"
		}
	}())

	// 空一行
	buffer += "\n"

	// method http://host:port?query protocol
	buffer += req.Method + " " + req.Uri + " " + req.Proto

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

// response 返回
func (server Server) response(ctx *gin.Context, r *result.Result) {
	// 结果导出请求
	if request2.IsCopyBody(ctx) {
		req := request2.Get(ctx)
		r.Request = &req
	}

	// 记录日志
	go func() {
		if r.Code == result.Ok.Code && server.httpAccessLog != nil {
			elapsed := middleware.GetElapsed(ctx)
			buffer := server.getResultBuffer(ctx, r, elapsed)
			server.httpAccessLog.Print(glog.Debug, buffer)
		} else if server.httpErrorLog != nil {
			elapsed := middleware.GetElapsed(ctx)
			buffer := server.getResultBuffer(ctx, r, elapsed)
			server.httpErrorLog.Print(glog.Error, buffer)
		}
	}()

	// 返回客户端
	ctx.AbortWithStatusJSON(http.StatusOK, r)
}
