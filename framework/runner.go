/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-12-01 03:59:13
 */

package framework

import (
	"framework/app"
	"framework/context"
)

// RunnerInterface 启动器
type RunnerInterface interface {
	// GetConfigJsonFile 获取配置文件
	GetConfigJsonFile() string
	// Mirage 初始化数据库
	Mirage()
	// GetHttpConfig 获取 http 配置
	GetHttpConfig() (groupName string, routerGroups []app.RouterGroup, port int)
}

// Run 启动
func Run(runnerInterface RunnerInterface) {
	// 加载配置文件
	if err := context.Init(runnerInterface.GetConfigJsonFile()); err != nil {
		panic(err.Error())
	}

	// 初始化数据库
	runnerInterface.Mirage()

	// 获取 app
	groupName, routerGroups, port := runnerInterface.GetHttpConfig()
	// 创建服务
	router := app.New()
	// 注册路由
	router.RouterGroup(groupName, routerGroups)
	// 启动服务
	router.Run(port)
}
