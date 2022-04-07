/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/21 00:47:21
 */

package user

import (
	"framework/app"
	"net/http"
	"service/middleware"
)

// GetRouter 获取路由对象
func GetRouter() app.Router {
	return app.Router{
		GroupName:  "/user",
		Middleware: middleware.Admin,
		RouterPaths: []app.RouterPath{
			{
				HttpMethod:   http.MethodPost,
				RelativePath: "/:id",
				HandlerFunc:  getUser,
			},
		},
	}
}
