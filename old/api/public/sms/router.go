/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/03 17:59:03
 */

package sms

import (
	"framework/app"
	"net/http"
)

// GetRouter 获取路由对象
func GetRouter() app.RouterGroup {
	return app.RouterGroup{
		GroupName:  "/sms",
		Middleware: nil,
		RouterPaths: []app.Router{
			{
				HttpMethod:   http.MethodPost,
				RelativePath: "/send",
				RouterFunc:   postSend,
			}, {
				HttpMethod:   http.MethodPost,
				RelativePath: "/verify",
				RouterFunc:   postVerify,
			},
		},
	}
}
