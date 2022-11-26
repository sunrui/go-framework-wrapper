/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/01
 */

package app

import (
	"framework/app/middleware"
	"framework/config"
	"github.com/gin-gonic/gin"
	"strconv"
)

// Server 服务对象
type Server struct {
	engine *gin.Engine // gin 对象
}

// New 创建新的服务对象
func New() *Server {
	engine := gin.New()

	// 注册中间件
	// 注册 404 回调
	engine.NoRoute(routerFunc(middleware.NotFound))

	// 注册 405 回调
	engine.HandleMethodNotAllowed = true
	engine.NoMethod(routerFunc(middleware.MethodNotAllowed))

	// 注册限流中间件
	rateLimit := middleware.NewRateLimit()
	engine.Use(routerFunc(rateLimit.Take))

	// 注册刷新令牌中间件
	engine.Use(middleware.Token)

	// 注册文档中间件
	// TODO 加入开关
	engine.GET("/doc/*any", middleware.Swagger)

	// 注册 body 中间件
	engine.Use(middleware.Body)

	// 注册异常中间件
	engine.Use(middleware.Recover())

	return &Server{
		engine: engine,
	}
}

// Middleware 中间件
func (server *Server) Middleware(handlerFunc gin.HandlerFunc) {
	server.engine.Use(handlerFunc)
}

// Router 路由对象
func (server *Server) Router(router RouterGroup) {
	registerRouter(server.engine, router)
}

// RouterGroup 路由对象组
func (server *Server) RouterGroup(groupName string, routers []RouterGroup) {
	for _, router := range routers {
		router.GroupName = groupName + router.GroupName
		registerRouter(server.engine, router)
	}
}

// Run 启动服务
func (server *Server) Run(port int) {
	if err := server.engine.Run(":" + strconv.Itoa(port)); err != nil {
		panic(err.Error())
	}
}

// 初始化
func init() {
	// 如果非调式环境注册 release 模式
	if !config.IsDev() {
		gin.SetMode(gin.ReleaseMode)
	}
}
