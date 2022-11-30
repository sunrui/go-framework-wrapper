/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/21 00:47:21
 */

package user

import (
	"framework/app"
	"medium/middleware"
	"net/http"
)

// GetRouter 获取路由
func GetRouter() app.RouterGroup {
	return app.RouterGroup{
		GroupName:  "/api-user",
		Middleware: middleware.Auth,
		Routers: []app.Router{
			{
				HttpMethod:   http.MethodPost,
				RelativePath: "/:id",
				RouterFunc:   getUser,
			},
		},
	}
}
