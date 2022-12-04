/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-12-04 09:25:52
 */

package server

import (
	"fmt"
	"framework/app/env"
	"framework/app/result"
	"framework/app/server/request"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 获取结果数据
func (server Server) getResultBuffer(ctx *gin.Context, r *result.Result) string {
	// 获取 request
	req := request.Get(ctx)

	// method http://host:port?query protocol
	buffer := req.Method + " " + req.Uri + " " + req.Proto

	// userId
	if payload, _, err := server.token.GetPayload(ctx); err == nil {
		buffer += " - userId(" + payload.UserId + ")"
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

// response 返回
func (server Server) response(ctx *gin.Context, r *result.Result) {
	// 结果导出请求
	if request.IsCopyBody(ctx) {
		req := request.Get(ctx)
		r.Request = &req
	}

	go func() {
		// 控制台日志
		if env.IsDev() {
			fmt.Println(r)
		}

		// 记录日志
		buffer := server.getResultBuffer(ctx, r)
		if r.Code == result.Ok.Code {
			if server.httpAccessLog != nil {
				server.httpAccessLog.Debugln(buffer)
			}
		} else {
			if server.httpErrorLog != nil {
				server.httpErrorLog.Errorln(buffer)
			}
		}
	}()

	// 返回客户端
	ctx.AbortWithStatusJSON(http.StatusOK, r)
}
