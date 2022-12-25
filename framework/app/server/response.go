/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-12-04 09:25:52
 */

package server

import (
	"framework/app/build"
	"framework/app/glog"
	"framework/app/result"
	"framework/app/server/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

// response 返回
func (server Server) response(ctx *gin.Context, r result.Result) {
	// 获取请求参数
	req := server.GetRequest(ctx)
	r.Request = &req

	h := glog.Http{
		Result: r,
		UserId: func(ctx *gin.Context) *string {
			if payload, _, err := server.token.GetPayload(ctx); err == nil {
				return &payload.UserId
			} else {
				return nil
			}
		}(ctx),
		Elapsed: middleware.GetElapsed(ctx),
	}

	// 记录日志
	if r.Code == result.Ok.Code && server.httpAccessLog != nil {
		server.httpAccessLog.PrintHttp(glog.DebugLevel, h)
	} else if server.httpErrorLog != nil {
		server.httpErrorLog.PrintHttp(glog.ErrorLevel, h)
	}

	// 根据用户设置是否打印用户请求参数，默认在测试环境中全部打印
	request := ctx.DefaultQuery("request", "0")
	if !build.IsDev() || request == "false" || request == "0" {
		r.Request = nil
	}

	// 设置 request = true || request = 1 强制开启请求参数
	if request == "true" || request == "1" {
		r.Request = &req
	}

	// 在正式环境中，将内部错误相关的文件行不会返回给用户
	if !build.IsDev() && r.Code == result.InternalError.Code {
		m := r.Data.(result.M)
		delete(m, "file")
	}

	// 返回客户端
	ctx.AbortWithStatusJSON(http.StatusOK, r)
}
