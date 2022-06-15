/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-06-15 15:22:35
 */

package _sample

import (
	"framework/app"
	"net/http"
)

// GetRouter 获取路由对象
func GetRouter() app.Router {
	return app.Router{
		GroupName:  "/sample",
		Middleware: nil,
		RouterPaths: []app.RouterPath{
			{
				HttpMethod:   http.MethodGet,
				RelativePath: "",
				HandlerFunc:  getSample,
			},
			{
				HttpMethod:   http.MethodPut,
				RelativePath: "",
				HandlerFunc:  putSample,
			},
		},
	}
}
