/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022/04/03 00:05:03
 */

package request

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

// Request 请求对象
type Request struct {
	Ip      string   `json:"ip"`      // ip 地址
	Proto   string   `json:"proto"`   // 协议
	Method  string   `json:"method"`  // 请求方式
	Uri     string   `json:"uri"`     // 访问地址
	Headers []string `json:"headers"` // http 首部
	Body    *string  `json:"body"`    // 请求体
}

// IsExport 是否导出
func IsExport(ctx *gin.Context) bool {
	const key, value = "request", "export"
	return ctx.Query(key) == value || ctx.GetHeader(key) == value
}

// GetRequest 获取请求对象
func GetRequest(ctx *gin.Context) Request {
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
		Body: func(ctx *gin.Context) *string {
			body, exists := ctx.Get("body")
			if exists {
				bodyString := body.(string)
				return &bodyString
			} else {
				return nil
			}
		}(ctx),
	}
}

// CopyBody 复制 body
func CopyBody(ctx *gin.Context) {
	if data, err := ctx.GetRawData(); err != nil {
		panic(err.Error())
	} else if len(data) != 0 {
		ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))
		ctx.Set("body", string(data))
	}
}

// validator 较验
// https://github.com/go-playground/validator/

// PageRequest 分页请求对象
type PageRequest struct {
	Page     int `json:"page" form:"page" validate:"required,gte=1,lte=9999"`       // 分页，从 1 开始
	PageSize int `json:"pageSize" form:"pageSize" validate:"required,gte=1,lte=99"` // 分页大小
}
