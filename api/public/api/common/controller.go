/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-12-05 03:27:54
 */

package common

import (
	"framework/server"
	"net/http"
)

type Controller struct {
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
