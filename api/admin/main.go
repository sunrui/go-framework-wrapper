/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022/04/07 22:38:07
 */
package main

import (
	"framework/app/server"
	"medium/service"
)

// @title   Medium 公用接口文档
// @version 1.0
// @host    127.0.0.1:8080
// @BasePath
func main() {
	// 初始化数据库
	service.Mirage()

	// 创建服务
	router := server.New(service.Ctx.Config.Server,
		service.Ctx.Log.HttpAccess,
		service.Ctx.Log.HttpError,
		service.Ctx.Token.Jwt)

	// 注册路由
	router.RouterGroup("/public", []server.RouterGroup{})

	port := 8080
	service.Ctx.Log.Service.Info("service start: http://127.0.0.1:%d", port)

	// 启动服务
	if err := router.Run(port); err != nil {
		panic(err.Error())
	}
}
