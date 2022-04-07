/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022/04/07 22:38:07
 */
package main

import (
	"api-public/rest/area"
	"api-public/rest/auth"
	"api-public/rest/common"
	"api-public/rest/sms"
	"framework/app"
)

// @title    Medium 公用接口文档
// @version  1.0
// @host     127.0.0.1:8080
// @BasePath
func main() {
	// 创建服务
	server := app.New()

	// 注册路由
	server.RouterGroup("/public", []app.Router{
		common.GetRouter(),
		area.GetRouter(),
		sms.GetRouter(),
		auth.GetRouter(),
	})

	// 启动服务
	server.Run(8080)
}
