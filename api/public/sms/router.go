/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/03 17:59:03
 */

package sms

import (
	"medium-server-go/framework/app"
	"net/http"
)

// 获取路由对象
func GetRouter() app.Router {
	return app.Router{
		GroupName:  "/sms",
		Middleware: nil,
		RouterPaths: []app.RouterPath{
			{
				HttpMethod:   http.MethodPost,
				RelativePath: "/code",
				HandlerFunc:  postCode,
			}, {
				HttpMethod:   http.MethodPost,
				RelativePath: "/verify",
				HandlerFunc:  postVerify,
			},
		},
	}
}
