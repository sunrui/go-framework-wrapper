/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/01
 */

package app

import (
	"github.com/gin-gonic/gin"
	_ "medium-server-go/docs"
	"medium-server-go/framework/config"
	"strconv"
)

// Server 服务对象
type Server struct {
	engine *gin.Engine // gin 对象
}

// New 创建新的服务对象
func New() *Server {
	engine := gin.Default()

	// 注册中间件
	registerMiddleware(engine)

	return &Server{
		engine: engine,
	}
}

// RouterGroup 路由对象组
func (server *Server) RouterGroup(groupName string, routers []Router) {
	for _, router := range routers {
		router.GroupName = groupName + router.GroupName
		registerRouter(server.engine, router)
	}
}

// Run 启动服务
func (server *Server) Run(port int) {
	err := server.engine.Run(":" + strconv.Itoa(port))
	if err != nil {
		panic(err.Error())
	}
}

// 初始化
func init() {
	// 如果非调式环境注册 release 模式
	if !config.IsDebugMode() {
		gin.SetMode(gin.ReleaseMode)
	}
}
