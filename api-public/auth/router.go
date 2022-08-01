/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/02
 */

package auth

import (
	"framework/app"
	"net/http"
)

// GetRouter 获取路由对象
func GetRouter() app.RouterGroup {
	return app.RouterGroup{
		GroupName:  "/auth",
		Middleware: nil,
		RouterPaths: []app.Router{
			{
				HttpMethod:   http.MethodPost,
				RelativePath: "/login/phone",
				RouterFunc:   postLoginByPhone,
			}, {
				HttpMethod:   http.MethodPost,
				RelativePath: "/login/wechat",
				RouterFunc:   postLoginByWechat,
			}, {
				HttpMethod:   http.MethodGet,
				RelativePath: "/token",
				RouterFunc:   getToken,
			}, {
				HttpMethod:   http.MethodPost,
				RelativePath: "/logout",
				RouterFunc:   postLogout,
			},
		},
	}
}
