/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2021/12/31
 */
package main

import (
	"medium-server-go/api/admin"
	"medium-server-go/api/public"
	"medium-server-go/api/user"
	"medium-server-go/framework/app"
)

// @title    Medium 接口文档
// @version  1.0
// @host     127.0.0.1:8080
// @BasePath
func main() {
	// 创建服务
	server := app.New()

	// 注册路由
	server.RouterGroup("/public", public.GetRouters())
	server.RouterGroup("/admin", admin.GetRouters())
	server.RouterGroup("/user", user.GetRouters())

	// 启动服务
	server.Run(8080)
}
