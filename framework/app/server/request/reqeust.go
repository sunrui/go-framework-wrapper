/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-10-15 11:20:20
 */

package request

import (
	"github.com/gin-gonic/gin"
)

// Request 请求
type Request struct {
	Ip      string   `json:"ip"`      // ip 地址
	Proto   string   `json:"proto"`   // 协议
	Method  string   `json:"method"`  // 请求方式
	Uri     string   `json:"uri"`     // 访问地址
	Headers []string `json:"headers"` // server 首部
	Body    *string  `json:"body"`    // 请求体
}

// Get 获取
func Get(ctx *gin.Context) Request {
	return Request{
		Ip:     ctx.ClientIP(),
		Proto:  ctx.Request.Proto,
		Method: ctx.Request.Method,
		Uri: func(ctx *gin.Context) string {
			if ctx.Request.TLS != nil {
				return "https://"
			} else {
				return "http://"
			}
		}(ctx) + ctx.Request.Host + ctx.Request.RequestURI,
		Headers: func(ctx *gin.Context) []string {
			headers := make([]string, 0)

			for key, value := range ctx.Request.Header {
				headers = append(headers, key+": "+value[0])
			}

			return headers
		}(ctx),
		Body: GetBody(ctx),
	}
}
