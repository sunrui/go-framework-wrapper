/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/21 00:47:21
 */

package user

import (
	"medium-server-go/framework/app"
	"medium-server-go/middleware"
	"net/http"
)

// 获取路由对象
func GetRouter() app.Router {
	return app.Router{
		GroupName:  "/user",
		Middleware: middleware.AuthMiddleware,
		RouterPaths: []app.RouterPath{
			{
				HttpMethod:   http.MethodPost,
				RelativePath: "/:id",
				HandlerFunc:  getUser,
			},
		},
	}
}
