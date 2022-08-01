/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/03 17:59:03
 */

package common

import (
	"framework/app"
	"net/http"
)

// GetRouter 获取路由对象
func GetRouter() app.RouterGroup {
	return app.RouterGroup{
		GroupName:  "/common",
		Middleware: nil,
		RouterPaths: []app.Router{
			{
				HttpMethod:   http.MethodGet,
				RelativePath: "/",
				RouterFunc:   getIndex,
			},
			{
				HttpMethod:   http.MethodGet,
				RelativePath: "/build",
				RouterFunc:   getBuild,
			},
			{
				HttpMethod:   http.MethodGet,
				RelativePath: "/result",
				RouterFunc:   getResult,
			},
		},
	}
}
