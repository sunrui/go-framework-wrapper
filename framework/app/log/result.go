/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-15 00:21:43
 */

package log

import (
	"fmt"
	"framework/app/request"
	"framework/result"
	"framework/token"
	"github.com/gin-gonic/gin"
	"log"
)

// 获取结果内容
func getResult(ctx *gin.Context, r result.Result[any]) string {
	// 获取 request 对象
	req := request.Get(ctx)

	// 等级
	var levelType LevelType
	if r.Code == result.OK {
		levelType = INFO
	} else {
		levelType = ERROR
	}

	var buffer string

	// 时间 - 等级 - IP
	buffer = " - " + string(levelType) + " - " + req.Ip

	// userId
	if userId := token.GetUserId(ctx); userId != nil {
		buffer += " - userId(" + *userId + ")"
	}

	switch levelType {
	case TRACE:
	case DEBUG:
	case INFO:
	case WARNING:
	case ERROR:
		buffer = fmt.Sprintf("\033[1;37;41m%s\033[0m", buffer)
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

// Write 写入
func Write(ctx *gin.Context, r result.Result[any]) {
	if GetLevelType() == NONE {
		return
	}

	stream := getResult(ctx, r)
	log.Println(stream)
	fmt.Println(stream)
}
