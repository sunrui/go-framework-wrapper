/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2021/12/31
 */
package main

import (
	api_admin "medium-server-go/api-admin"
	api_public "medium-server-go/api-public"
	api_user "medium-server-go/api-user"
	"medium-server-go/framework/app"
)

func main() {
	// 创建服务
	server := app.New()

	// 注册路由
	api_public.Register(server)
	api_user.Register(server)
	api_admin.Register(server)

	// 启动服务
	server.Run(8080)
}
