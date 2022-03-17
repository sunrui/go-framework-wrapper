/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/03 17:59:03
 */

package api_sms

import (
	"medium-server-go/common/app-gin"
	"net/http"
)

// 获取路由对象
func GetRouter() app_gin.Router {
	return app_gin.Router{
		GroupName: "/sms",
		RoleType:  app_gin.RolePublic,
		RouterPaths: []app_gin.RouterPath{
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
