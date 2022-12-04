/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-29 22:24:46
 */

package sms

import (
	"framework/app"
	"net/http"
)

// GetRouter 获取路由
func GetRouter() server.RouterGroup {
	return server.RouterGroup{
		GroupName:  "/sms",
		Middleware: nil,
		Routers: []server.Router{
			{
				HttpMethod:   http.MethodPost,
				RelativePath: "send",
				RouterFunc:   postSend,
			},
			{
				HttpMethod:   http.MethodPost,
				RelativePath: "verify",
				RouterFunc:   postVerify,
			},
		},
	}
}
