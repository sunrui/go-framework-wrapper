/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-12-16 20:32:03
 */

package server

import (
	"framework/app/result"
	"framework/app/server/body"
	"github.com/gin-gonic/gin"
)

// GetRequest 获取
func (server Server) GetRequest(ctx *gin.Context) result.Request {
	return result.Request{
		Ip:     ctx.ClientIP(),
		Method: ctx.Request.Method,
		Uri: func(ctx *gin.Context) string {
			if ctx.Request.TLS != nil {
				return "https://"
			} else {
				return "http://"
			}
		}(ctx) + ctx.Request.Host + ctx.Request.RequestURI,
		Header: ctx.Request.Header,
		Body:   body.GetBody(ctx),
	}
}
