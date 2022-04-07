/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022/04/07 22:38:07
 */
package main

import (
	"api-admin/rest/user"
	"framework/app"
)

// @title    Medium 接口文档
// @version  1.0
// @host     127.0.0.1:8080
// @BasePath
func main() {
	// 创建服务
	server := app.New()

	// 注册路由
	server.RouterGroup("/admin", []app.Router{
		user.GetRouter(),
	})

	// 注册演示路由

	// 启动服务
	server.Run(8080)
}
