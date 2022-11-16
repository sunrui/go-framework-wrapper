/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-16 21:44:52
 */

package log

import (
	"framework/app/request"
	"framework/result"
	"framework/token"
	"github.com/gin-gonic/gin"
)

// 获取结果内容
func getBuffer(ctx *gin.Context, r result.Result[any]) string {
	// 获取 request 对象
	req := request.Get(ctx)

	var buffer string

	// ip
	buffer = req.Ip

	// userId
	if userId := token.GetUserId(); userId != nil {
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

	return buffer
}

// WriteResult 写入文件
func WriteResult(ctx *gin.Context, r result.Result[any]) {
	stream := getBuffer(ctx, r)

	if r.Code == result.OK {
		logger.Debug(stream)
	} else {
		logger.Error(stream)
	}
}
