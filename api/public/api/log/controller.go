/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-12-05 04:10:16
 */

package log

import (
	"framework/app/server"
	"medium/service"
	"medium/service/log"
	"net/http"
)

// Controller 控制器
type Controller struct {
	ctx               *service.Context      // 上下文
	logHttpRepository log.LogHttpRepository // LogHttp 日志访问
}

// NewController 创建控制器
func NewController(ctx *service.Context) Controller {
	return Controller{
		ctx:               ctx,
		logHttpRepository: log.NewLogHttpRepository(ctx.Mysql),
	}
}

// GetRouter 获取路由
func (controller Controller) GetRouter() server.RouterGroup {
	return server.RouterGroup{
		GroupName:  "/log",
		Middleware: nil,
		Routers: []server.Router{
			{
				HttpMethod:   http.MethodGet,
				RelativePath: "/",
				RouterFunc:   controller.getIndex,
			},
		},
	}
}
