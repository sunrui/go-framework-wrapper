/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-14 23:35:55
 */

package request

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io"
)

// BodyTag 上下文 body 标记
const BodyTag = "BodyTag"

// CopyBody 复制 body
func CopyBody(ctx *gin.Context) {
	if data, err := ctx.GetRawData(); err != nil {
		panic(err.Error())
	} else if len(data) != 0 {
		ctx.Request.Body = io.NopCloser(bytes.NewBuffer(data))
		ctx.Set(BodyTag, string(data))
	}
}

// GetBody 获取 body
func GetBody(ctx *gin.Context) *string {
	body, exists := ctx.Get(BodyTag)
	if exists {
		bodyString := body.(string)
		return &bodyString
	} else {
		return nil
	}
}
