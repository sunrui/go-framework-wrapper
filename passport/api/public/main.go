/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022/04/07 22:38:07
 */
package main

import (
	"framework/app"
	"framework/context"
	"passport/service"
	"path/filepath"
	"public/api/common"
	"public/api/user"
	"runtime"
)

// 初始化上下文
func initContext() {
	_, file, _, _ := runtime.Caller(0)
	path := filepath.Dir(file)

	if err := context.Init(path + "/../../config.json"); err != nil {
		panic(err.Error())
	}
}

// 开启服务
func startServer() {
	// 创建服务
	server := app.New()

	// 注册路由
	server.RouterGroup("/public", []app.RouterGroup{
		common.GetRouter(),
		user.GetRouter(),
	})

	// 启动服务
	server.Run(8080)
}

// @title   Medium 公用接口文档
// @version 1.0
// @host    127.0.0.1:8080
// @BasePath
func main() {
	// 初始化上下文
	initContext()

	// 数据库初始化
	service.Mirage()

	// 启动服务
	startServer()
}
