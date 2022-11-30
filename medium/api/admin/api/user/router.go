/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/21 00:47:21
 */

package user

import (
	"framework/http"
	"medium/middleware"
	"net/http"
)

// GetRouter 获取路由
func GetRouter() http.RouterGroup {
	return http.RouterGroup{
		GroupName:  "/api-user",
		Middleware: middleware.Admin,
		Routers: []http.Router{
			{
				HttpMethod:   http.MethodPost,
				RelativePath: "/:id",
				RouterFunc:   getUser,
			},
		},
	}
}
