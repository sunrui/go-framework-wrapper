/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-16 21:44:52
 */

package log

import (
	"framework/app/request"
	"framework/app/token"
	"framework/config"
	"framework/result"
	"github.com/gin-gonic/gin"
)

// WriteResult 获取结果内容
func WriteResult(ctx *gin.Context, r *result.Result) {
	if !config.Inst().Log.Enable {
		return
	}

	// 获取 request 对象
	req := request.Get(ctx)

	var buffer string

	// ip
	buffer = req.Ip

	// userId
	if userId := token.GetUserId(ctx); userId != nil {
		buffer += " - userId(" + *userId + ")"
	}

	// 换行
	buffer += "\n\n"

	// method http://host:port?query protocol
	buffer += req.Method + " " + req.Uri + " " + req.Proto + "\n"

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

	if r.Code == result.Ok.Code {
		Inst.Debug(buffer)
	} else {
		Inst.Error(buffer)
	}
}
