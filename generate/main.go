/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022/04/07 23:29:07
 */

package main

import (
	"framework/app"
	admin "generate/api-admin/template"
	user "generate/api-user/template"
)

// @title    Medium 模板接口文档
// @version  1.0
// @host     127.0.0.1:8080
// @BasePath
func main() {
	// 创建服务
	server := app.New()

	// 注册路由
	server.RouterGroup("/", []app.Router{
		user.GetRouter(),
		admin.GetRouter(),
	})

	// 启动服务
	server.Run(8080)
}