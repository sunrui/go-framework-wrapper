/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-29 22:24:46
 */

package sms

import (
	"framework/http"
	"net/http"
)

// GetRouter 获取路由
func GetRouter() http.RouterGroup {
	return http.RouterGroup{
		GroupName:  "/sms",
		Middleware: nil,
		Routers: []http.Router{
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
