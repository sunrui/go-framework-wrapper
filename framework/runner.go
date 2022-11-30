/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-12-01 03:59:13
 */

package framework

import (
	"framework/context"
	"framework/http"
)

// RunnerInterface 启动器
type RunnerInterface interface {
	// GetConfigJson 获取配置文件
	GetConfigJson() string
	// Mirage 初始化数据库
	Mirage()
	// GetHttp 获取 http
	GetHttp() (groupName string, routerGroups []http.RouterGroup, port int)
}

// Run 启动
func Run(runnerInterface RunnerInterface) {
	// 加载配置文件
	if err := context.Init(runnerInterface.GetConfigJson()); err != nil {
		panic(err.Error())
	}

	// 初始化数据库
	runnerInterface.Mirage()

	// 获取 http
	groupName, routerGroups, port := runnerInterface.GetHttp()
	// 创建服务
	router := http.New()
	// 注册路由
	router.RouterGroup(groupName, routerGroups)
	// 启动服务
	router.Run(port)
}
