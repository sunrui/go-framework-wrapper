/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022/04/07 22:38:07
 */
package main

import (
	"framework"
	"framework/app"
	"medium/service"
	"path/filepath"
	"runtime"
)

// Runner 启动器实例
type Runner struct {
}

// GetConfigJsonFile 获取配置文件
func (Runner) GetConfigJsonFile() string {
	_, file, _, _ := runtime.Caller(0)
	path := filepath.Dir(file)

	return path + "/../../config.json"
}

// Mirage 初始化数据库
func (Runner) Mirage() {
	service.Mirage()
}

// GetHttpConfig 获取 http 配置
func (Runner) GetHttpConfig() (groupName string, routerGroups []app.RouterGroup, port int) {
	return "/public", []app.RouterGroup{
		//common.GetRouter(),
		//user.GetRouter(),
	}, 8080
}

// @title   Medium 公用接口文档
// @version 1.0
// @host    127.0.0.1:8080
// @BasePath
func main() {
	framework.Run(Runner{})
}
