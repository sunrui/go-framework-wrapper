/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-26 13:37:48
 */

package user

import (
	"framework/app"
	"net/http"
)

// GetRouter 获取路由
func GetRouter() app.RouterGroup {
	return app.RouterGroup{
		GroupName:  "/user",
		Middleware: nil,
		Routers: []app.Router{
			{
				HttpMethod:   http.MethodPost,
				RelativePath: "login/phone",
				RouterFunc:   postLoginByPhone,
			},
		},
	}
}
