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
func (server Server) response(ctx *gin.Context, r *result.Result) {
	// 结果导出请求
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

	if r.Code == result.Ok.Code && server.httpAccessLog != nil {
		server.httpAccessLog.PrintHttp(glog.DebugLevel, h)
	} else if server.httpErrorLog != nil {
		server.httpErrorLog.PrintHttp(glog.ErrorLevel, h)
	}

	// 返回客户端的时候不返回  request
	dump := ctx.DefaultQuery("request", "")
	if !build.IsDev() || dump == "false" || dump == "0" {
		r.Request = nil
	}

	// 返回客户端
	ctx.AbortWithStatusJSON(http.StatusOK, r)
}
