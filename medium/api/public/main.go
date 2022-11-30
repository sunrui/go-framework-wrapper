/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022/04/07 22:38:07
 */
package main

import (
	"framework"
	"framework/http"
	"medium/service"
	"path/filepath"
	"runtime"
)

// Runner 启动器实例
type Runner struct {
}

// GetConfigJson 获取配置文件
func (Runner) GetConfigJson() string {
	_, file, _, _ := runtime.Caller(0)
	path := filepath.Dir(file)

	return path + "/../../config.json"
}

// Mirage 初始化数据库
func (Runner) Mirage() {
	service.Mirage()
}

// GetHttp 获取 http
func (Runner) GetHttp() (groupName string, routerGroups []http.RouterGroup, port int) {
	return "/public", []http.RouterGroup{
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
