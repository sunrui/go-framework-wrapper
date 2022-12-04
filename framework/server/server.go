/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/01
 */

package server

import (
	"framework/app/env"
	"framework/app/log"
	"framework/app/token"
	middleware2 "framework/server/middleware"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"strconv"
)

// Config 配置
type Config struct {
	RateLimitCapacity int64 `json:"rateLimitCapacity"` // 令牌桶容量
	RateLimitQuantum  int64 `json:"rateLimitQuantum"`  // 每隔多少秒
	EnableDoc         bool  `json:"enableDoc"`         // 启用文档
}

// Server 服务
type Server struct {
	engine        *gin.Engine  // gin config
	httpAccessLog *log.Log     // api 访问日志
	httpErrorLog  *log.Log     // api 错误日志
	token         *token.Token // 令牌
}

// New 创建服务
func New(config Config, httpAccessLog *log.Log, httpErrorLog *log.Log, token *token.Token) *Server {
	engine := gin.New()

	server := &Server{
		engine:        engine,
		httpAccessLog: httpAccessLog,
		httpErrorLog:  httpErrorLog,
		token:         token,
	}

	if env.IsDev() {
		gin.SetMode(gin.DebugMode)
		gin.DefaultWriter = io.MultiWriter(httpAccessLog.Out, os.Stdout)
	} else {
		gin.SetMode(gin.ReleaseMode)
		gin.DefaultWriter = io.MultiWriter(httpAccessLog.Out)
	}

	// 注册 404 回调
	engine.NoRoute(server.routerFunc(middleware2.NotFound))

	// 注册 405 回调
	engine.HandleMethodNotAllowed = true
	engine.NoMethod(server.routerFunc(middleware2.MethodNotAllowed))

	// 注册限流中间件
	rateLimit := middleware2.NewRateLimit(config.RateLimitCapacity, config.RateLimitQuantum)
	engine.Use(server.routerFunc(rateLimit.Take))

	// 注册令牌中间件
	engine.Use(middleware2.Token)

	// 注册文档中间件
	if config.EnableDoc {
		// 注册 body 中间件
		engine.Use(middleware2.Body)

		engine.GET("/doc/*any", middleware2.Swagger)
	}

	// 注册异常中间件
	engine.Use(server.routerFunc(middleware2.Recover))

	return &Server{
		engine:        engine,
		httpAccessLog: httpAccessLog,
		httpErrorLog:  httpErrorLog,
		token:         token,
	}
}

// Middleware 中间件
func (server Server) Middleware(handlerFunc gin.HandlerFunc) {
	server.engine.Use(handlerFunc)
}

// Router 路由
func (server Server) Router(routerGroup RouterGroup) {
	server.registerRouter(server.engine, routerGroup)
}

// RouterGroup 路由组
func (server Server) RouterGroup(groupName string, routerGroups []RouterGroup) {
	for _, router := range routerGroups {
		router.GroupName = groupName + router.GroupName
		server.registerRouter(server.engine, router)
	}
}

// Run 启动服务
func (server Server) Run(port int) error {
	return server.engine.Run(":" + strconv.Itoa(port))
}
