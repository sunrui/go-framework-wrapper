/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/03 17:59:03
 */

package common

import (
	"medium-server-go/framework/app"
	"net/http"
)

// GetRouter 获取路由对象
func GetRouter() app.Router {
	return app.Router{
		GroupName:  "/common",
		Middleware: nil,
		RouterPaths: []app.RouterPath{
			{
				HttpMethod:   http.MethodGet,
				RelativePath: "/build",
				HandlerFunc:  getBuild,
			},
			{
				HttpMethod:   http.MethodGet,
				RelativePath: "/result",
				HandlerFunc:  getResult,
			},
		},
	}
}
