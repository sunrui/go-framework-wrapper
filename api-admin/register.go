/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/25 22:29:25
 */

package api_admin

import (
	"medium-server-go/api-admin/api-user"
	"medium-server-go/framework/app"
)

// 注册路由
func Register(server *app.Server) {
	for _, router := range []app.Router{
		api_user.GetRouter(),
	} {
		router.GroupName = "/admin" + router.GroupName
		server.RegisterRouter(router)
	}
}
