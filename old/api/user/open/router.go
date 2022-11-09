/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/31 21:00:31
 */

package open

import (
	"framework/app"
	"middleware"
	"net/http"
)

// GetRouter 获取路由对象
func GetRouter() app.RouterGroup {
	return app.RouterGroup{
		GroupName:  "/open",
		Middleware: middleware.Auth,
		RouterPaths: []app.Router{
			{
				HttpMethod:   http.MethodGet,
				RelativePath: "",
				RouterFunc:   getOpen,
			},
			{
				HttpMethod:   http.MethodPost,
				RelativePath: "",
				RouterFunc:   postOpen,
			},
		},
	}
}
