/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/31 21:00:31
 */

package open

import (
	"medium-server-go/framework/app"
	"medium-server-go/middleware"
	"net/http"
)

// 获取路由对象
func GetRouter() app.Router {
	return app.Router{
		GroupName:  "/open",
		Middleware: middleware.AuthMiddleware,
		RouterPaths: []app.RouterPath{
			{
				HttpMethod:   http.MethodGet,
				RelativePath: "",
				HandlerFunc:  getOpen,
			},
			{
				HttpMethod:   http.MethodPost,
				RelativePath: "",
				HandlerFunc:  postOpen,
			},
		},
	}
}
