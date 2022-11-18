/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-18 22:21:00
 */

package main

import (
	"framework/app"
	"passport/api/public"
)

// @title   Medium 公用接口文档
// @version 1.0
// @host    127.0.0.1:8080
// @BasePath
func main() {
	// 创建服务
	server := app.New()

	// 注册路由
	server.RouterGroup("/public", []app.RouterGroup{
		public.GetRouter(),
	})

	// 启动服务
	server.Run(8080)
}
