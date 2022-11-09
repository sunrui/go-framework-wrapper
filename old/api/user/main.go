/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022/04/07 22:50:07
 */
package main

import (
	"framework/app"
	"user/open"
	"user/user"
)

// @title    Medium 用户接口文档
// @version  1.0
// @host     127.0.0.1:8081
// @BasePath
func main() {
	// 创建服务
	server := app.New()

	// 注册路由
	server.RouterGroup("/user", []app.RouterGroup{
		open.GetRouter(),
		user.GetRouter(),
	})

	// 启动服务
	server.Run(8081)
}
