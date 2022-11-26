/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022/04/07 22:38:07
 */
package main

import (
	"admin/api/user"
	"framework/app"
)

// @title   Medium 后台接口文档
// @version 1.0
// @host    127.0.0.1:8082
// @BasePath
func main() {
	// 创建服务
	server := app.New()

	// 注册路由
	server.RouterGroup("/api-admin", []app.RouterGroup{
		user.GetRouter(),
	})

	// 启动服务
	server.Run(8082)
}
