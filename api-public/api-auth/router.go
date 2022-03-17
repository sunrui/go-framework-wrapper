/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/02
 */
package api_auth

import (
	"medium-server-go/common/app-gin"
	"net/http"
)

// 获取路由对象
func GetRouter() app_gin.Router {
	return app_gin.Router{
		GroupName: "/auth",
		RoleType:  app_gin.RolePublic,
		RouterPaths: []app_gin.RouterPath{
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
