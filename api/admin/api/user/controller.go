/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-12-24 16:46:03
 */

package user

import (
	"framework/app/server"
	"medium/middleware"
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
		GroupName:  "/server-user",
		Middleware: middleware.Admin,
		Routers: []server.Router{
			{
				HttpMethod:   http.MethodPost,
				RelativePath: "/:id",
				RouterFunc:   getUser,
			},
		},
	}
}
