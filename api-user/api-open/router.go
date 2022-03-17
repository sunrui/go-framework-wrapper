/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/31 21:00:31
 */

package api_open

import (
	"medium-server-go/framework/app-gin"
	"net/http"
)

// 获取路由对象
func GetRouter() app_gin.Router {
	return app_gin.Router{
		GroupName: "/open",
		RoleType:  app_gin.RoleAuth,
		RouterPaths: []app_gin.RouterPath{
			{
				HttpMethod:   http.MethodGet,
				RelativePath: "",
				HandlerFunc:  getOpen,
			},
			{
				HttpMethod:   http.MethodPost,
				RelativePath: "",
				HandlerFunc:  postOpen,
			},
		},
	}
}
