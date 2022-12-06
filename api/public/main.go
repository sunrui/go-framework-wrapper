/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022/04/07 22:38:07
 */
package main

import (
	"framework/app/server"
	"medium/service"
	"public/api/common"
	"public/api/log"
)

// @title   Medium 公用接口文档
// @version 1.0
// @host    127.0.0.1:8080
// @BasePath
func main() {
	var context *service.Context
	var err error

	// 创建上下文
	if context, err = service.NewContext("config.json"); err != nil {
		panic(err.Error())
	}

	// 创建服务
	router := server.New(context.Config.Server,
		context.Log.HttpAccess,
		context.Log.HttpError,
		context.Token.Jwt)

	// 注册路由
	router.RouterGroup("/public", []server.RouterGroup{
		common.GetRouter(),
		log.NewController(context.Mysql).GetRouter(),
	})

	// 端口
	const port = 8080
	context.Log.Service.Info("service start: http://127.0.0.1:%d", port)

	// 启动服务
	if err = router.Run(port); err != nil {
		panic(err.Error())
	}
}
