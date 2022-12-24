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
	var ctx *service.Context
	var err error

	// 创建上下文
	if ctx, err = service.NewContext(); err != nil {
		panic(err.Error())
	}

	// 创建服务
	svr := server.New(ctx.Config.Server,
		ctx.Log.HttpAccess,
		ctx.Log.HttpError,
		ctx.Token.Jwt)

	// 注册路由
	svr.RouterGroup("/public", []server.RouterGroup{
		common.NewController(ctx).GetRouter(),
		log.NewController(ctx).GetRouter(),
	})

	// 启动服务
	if err = svr.Run(8080); err != nil {
		panic(err.Error())
	}
}
