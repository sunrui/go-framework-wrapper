/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-12-04 09:25:52
 */

package server

import (
	"framework/app/glog"
	"framework/app/result"
	"framework/server/middleware"
	"framework/server/request"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 获取结果数据
func (server Server) getResultBuffer(ctx *gin.Context, r *result.Result) string {
	// 获取 request
	req := request.Get(ctx)

	// method http://host:port?query protocol
	buffer := req.Method + " " + req.Uri + " " + req.Proto

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

// 获取格式化
func (server Server) getFormat(ctx *gin.Context, r *result.Result) glog.Format {
	return glog.Format{
		Request: r.Request,
		Level: func() glog.Level {
			if r.Code == result.Ok.Code {
				return glog.DebugLevel
			} else {
				return glog.ErrorLevel
			}
		}(),
		Message: func() string {
			return server.getResultBuffer(ctx, r)
		}(),
		Elapsed: r.Elapsed,
		UserId: func() *string {
			if payload, _, err := server.token.GetPayload(ctx); err == nil {
				return &payload.UserId
			} else {
				return nil
			}
		}(),
	}
}

// response 返回
func (server Server) response(ctx *gin.Context, r *result.Result) {
	// 结果导出请求
	if request.IsCopyBody(ctx) {
		req := request.Get(ctx)
		r.Request = &req
	}

	// 记录日志
	go func() {
		r.Elapsed = middleware.GetElapsed(ctx)

		if r.Code == result.Ok.Code && server.httpAccessLog != nil {
			server.httpAccessLog.PrintMessage(server.getFormat(ctx, r))
		} else if server.httpErrorLog != nil {
			server.httpErrorLog.PrintMessage(server.getFormat(ctx, r))
		}
	}()

	// 返回客户端
	ctx.AbortWithStatusJSON(http.StatusOK, r)
}
