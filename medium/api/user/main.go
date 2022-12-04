/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022/04/07 22:38:07
 */
package main

import (
	"framework/app/env"
	"framework/context"
	"framework/server"
	"medium/service"
	"path/filepath"
	"runtime"
)

var Context *context.Context

// @title   Medium 公用接口文档
// @version 1.0
// @host    127.0.0.1:8080
// @BasePath
func main() {
	var err error

	_, file, _, _ := runtime.Caller(0)
	path := filepath.Dir(file)

	if Context, err = context.New(path+"/../../config.json", env.IsDev()); err != nil {
		panic(err.Error())
	}

	service.Mirage()

	// 创建服务
	router := server.New(Context.Config.Server, Context.HttpAccessLog, Context.HttpErrorLog, Context.JwtToken, env.IsDev())

	// 注册路由
	router.RouterGroup("/public", []server.RouterGroup{})

	// 启动服务
	if err = router.Run(8080); err != nil {
		panic(err.Error())
	}
}
