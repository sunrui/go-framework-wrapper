/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-29 22:24:46
 */

package sms

import (
	"framework/app/server"
	"medium/service"
	"net/http"
)

// Controller 控制器
type Controller struct {
	ctx *service.Context // 上下文
}

// NewController 创建控制器
func NewController(ctx *service.Context) Controller {
	return Controller{
		ctx: ctx,
	}
}

// GetRouter 获取路由
func (controller Controller) GetRouter() server.RouterGroup {
	return server.RouterGroup{
		GroupName:  "/sms",
		Middleware: nil,
		Routers: []server.Router{
			{
				HttpMethod:   http.MethodPost,
				RelativePath: "send",
				RouterFunc:   postSend,
			},
			{
				HttpMethod:   http.MethodPost,
				RelativePath: "verify",
				RouterFunc:   postVerify,
			},
		},
	}
}
