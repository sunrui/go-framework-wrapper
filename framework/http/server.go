/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/01
 */

package http

import (
	"framework/config"
	"framework/context"
	"framework/http/middleware"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"strconv"
)

// Server 服务
type Server struct {
	engine *gin.Engine // gin config
}

// New 创建服务
func New() *Server {
	engine := gin.New()

	if config.IsDev() {
		gin.SetMode(gin.DebugMode)
		gin.DefaultWriter = io.MultiWriter(context.Log.HttpAccess.Out, os.Stdout)
	} else {
		gin.SetMode(gin.ReleaseMode)
		gin.DefaultWriter = io.MultiWriter(context.Log.HttpAccess.Out)
	}

	// 注册 404 回调
	engine.NoRoute(routerFunc(middleware.NotFound))

	// 注册 405 回调
	engine.HandleMethodNotAllowed = true
	engine.NoMethod(routerFunc(middleware.MethodNotAllowed))

	// 注册限流中间件
	engine.Use(routerFunc(middleware.NewRateLimit().Take))

	// 注册刷新令牌中间件
	engine.Use(middleware.Token)

	// 注册文档中间件
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

// Router 路由
func (server *Server) Router(routerGroup RouterGroup) {
	registerRouter(server.engine, routerGroup)
}

// RouterGroup 路由组
func (server *Server) RouterGroup(groupName string, routerGroups []RouterGroup) {
	for _, router := range routerGroups {
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
