/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/21 00:47:21
 */

package user

import (
	"medium/middleware"
	"net/http"
)

// GetRouter 获取路由
func GetRouter() server.RouterGroup {
	return server.RouterGroup{
		GroupName:  "/server-user",
		Middleware: middleware.Auth,
		Routers: []server.Router{
			{
				HttpMethod:   http.MethodPost,
				RelativePath: "/:id",
				RouterFunc:   getUser,
			},
		},
	}
}
