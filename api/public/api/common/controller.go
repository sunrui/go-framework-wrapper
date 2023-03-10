/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/03 17:59:03
 */

package common

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
		GroupName:  "/common",
		Middleware: nil,
		Routers: []server.Router{
			{
				HttpMethod:   http.MethodGet,
				RelativePath: "/error",
				RouterFunc:   getError,
			},
			{
				HttpMethod:   http.MethodGet,
				RelativePath: "/build",
				RouterFunc:   getBuild,
			},
			{
				HttpMethod:   http.MethodPost,
				RelativePath: "/send",
				RouterFunc:   postSend,
			},
		},
	}
}
