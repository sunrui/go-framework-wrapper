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
	"github.com/sirupsen/logrus"
)

// 获取结果内容
func getBuffer(ctx *gin.Context, r result.Result[any]) string {
	// 获取 request 对象
	req := request.Get(ctx)

	// 等级
	var logLevel logrus.Level
	if r.Code == result.OK {
		logLevel = logrus.DebugLevel
	} else {
		logLevel = logrus.ErrorLevel
	}

	var buffer string

	// 时间 - 等级 - IP
	buffer = " - " + logLevel.String() + " - " + req.Ip

	// userId
	if userId := token.GetUserId(); userId != nil {
		buffer += " - userId(" + *userId + ")"
	}

	// 换行
	buffer += "\n"

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

// WriteLog 写入文件
func WriteLog(ctx *gin.Context, r result.Result[any]) {
	stream := getBuffer(ctx, r)
	logger.Println(stream)
}
