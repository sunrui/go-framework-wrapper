/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/01
 */

package server

import (
	"framework/app/glog"
	"framework/app/server/middleware"
	"framework/app/token"
	"github.com/gin-gonic/gin"
	"strconv"
)

// Config 配置
type Config struct {
	RateLimitCapacity float64 `json:"rateLimitCapacity"` // 令牌桶容量
	EnableDoc         bool    `json:"enableDoc"`         // 启用文档
}

// Server 服务
type Server struct {
	engine        *gin.Engine  // gin config
	httpAccessLog *glog.GLog   // api 访问日志
	httpErrorLog  *glog.GLog   // api 错误日志
	token         *token.Token // 令牌
}

// New 创建服务
func New(config Config, httpAccessLog *glog.GLog, httpErrorLog *glog.GLog, token *token.Token) *Server {
	engine := gin.New()

	server := &Server{
		engine:        engine,
		httpAccessLog: httpAccessLog,
		httpErrorLog:  httpErrorLog,
		token:         token,
	}

	// 注册耗时中件间
	engine.Use(middleware.Elapsed)

	// 注册异常中间件
	engine.Use(func(ctx *gin.Context) {
		if r := middleware.Recover(ctx); r != nil {
			server.response(ctx, *r)
		}
	})

	// 注册 404 回调
	engine.NoRoute(middleware.NotFound)

	// 注册 405 回调
	engine.HandleMethodNotAllowed = true
	engine.NoMethod(middleware.MethodNotAllowed)

	// // 注册限流中间件
	rateLimit := middleware.NewRateLimit(config.RateLimitCapacity)
	engine.Use(rateLimit.Take)

	// 注册令牌中间件
	engine.Use(middleware.Token)

	// 注册文档中间件
	if config.EnableDoc {
		engine.GET("/doc/*any", middleware.Swagger)
	}

	// 注册 body 中间件
	engine.Use(middleware.Body)

	return server
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
