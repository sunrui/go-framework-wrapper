/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-26 20:39:15
 */

package response

import (
	"framework/app/request"
	"framework/context"
	"framework/result"
	"github.com/gin-gonic/gin"
)

// 获取结果数据
func getResultString(ctx *gin.Context, r *result.Result) string {
	// 获取 request
	req := request.Get(ctx)

	// method http://host:port?query protocol
	buffer := req.Method + " " + req.Uri + " " + req.Proto

	// userId
	if payload, _, err := context.Token.GetPayload(ctx); err != nil {
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
