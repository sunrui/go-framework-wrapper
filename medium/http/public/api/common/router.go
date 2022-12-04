/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/03 17:59:03
 */

package common

import (
	"framework/app/server"
	"net/http"
)

// GetRouter 获取路由
func GetRouter() server.RouterGroup {
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
