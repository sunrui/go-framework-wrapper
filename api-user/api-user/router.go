/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/21 00:47:21
 */

package api_user

import (
	"medium-server-go/common/app-gin"
	"net/http"
)

// 获取路由对象
func GetRouter() app_gin.Router {
	return app_gin.Router{
		GroupName: "/user",
		RoleType:  app_gin.RoleAuth,
		RouterPaths: []app_gin.RouterPath{
			{
				HttpMethod:   http.MethodPost,
				RelativePath: "/:id",
				HandlerFunc:  getUser,
			},
		},
	}
}
