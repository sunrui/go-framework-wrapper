/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/02
 */
package auth

import (
	"medium-server-go/framework/app"
	"net/http"
)

// 获取路由对象
func GetRouter() app.Router {
	return app.Router{
		GroupName:  "/auth",
		Middleware: nil,
		RouterPaths: []app.RouterPath{
			{
				HttpMethod:   http.MethodPost,
				RelativePath: "/login/phone",
				HandlerFunc:  postLoginByPhone,
			}, {
				HttpMethod:   http.MethodPost,
				RelativePath: "/login/wechat",
				HandlerFunc:  postLoginByWechat,
			}, {
				HttpMethod:   http.MethodGet,
				RelativePath: "/token",
				HandlerFunc:  getToken,
			}, {
				HttpMethod:   http.MethodPost,
				RelativePath: "/logout",
				HandlerFunc:  postLogout,
			},
		},
	}
}
