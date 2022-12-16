/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-12-16 21:21:55
 */

package body

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io"
)

// 上下文 body 标记
const bodyTag = "bodyTag"

// CopyBody 复制 body
func CopyBody(ctx *gin.Context) {
	if data, err := ctx.GetRawData(); err != nil {
		panic(err.Error())
	} else if len(data) != 0 {
		ctx.Request.Body = io.NopCloser(bytes.NewBuffer(data))
		ctx.Set(bodyTag, string(data))
	}
}

// GetBody 获取 body
func GetBody(ctx *gin.Context) *string {
	if body, ok := ctx.Get(bodyTag); ok {
		bodyString := body.(string)
		return &bodyString
	} else {
		return nil
	}
}
