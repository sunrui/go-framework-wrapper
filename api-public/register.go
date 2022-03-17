/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/25 22:29:25
 */

package api_public

import (
	"medium-server-go/api-public/api-area"
	"medium-server-go/api-public/api-auth"
	"medium-server-go/api-public/api-sms"
	"medium-server-go/framework/app-gin"
)

// 注册路由
func Register(server *app_gin.Server) {
	for _, router := range []app_gin.Router{
		api_area.GetRouter(),
		api_sms.GetRouter(),
		api_auth.GetRouter(),
	} {
		server.RegisterRouter(router)
	}
}
