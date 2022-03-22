/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/01
 */
package app

import (
	"github.com/gin-gonic/gin"
	"medium-server-go/framework/config"
	"medium-server-go/framework/result"
	"net/http"
	"strconv"
	"time"
)

// 服务对象
type Server struct {
	engine *gin.Engine
}

// 初始化
func init() {
	// 如果非调式环境注册 release 模式
	if !config.IsDebugMode() {
		gin.SetMode(gin.ReleaseMode)
	}
}

// 创建新的服务对象
func New() *Server {
	engine := gin.Default()

	// 注册 404 回调
	engine.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusBadRequest, result.NotFound.WithKeyPair("uri", ctx.Request.URL.RequestURI()))
	})

	// 注册 405 回调
	{
		engine.HandleMethodNotAllowed = true
		engine.NoMethod(func(ctx *gin.Context) {
			ctx.JSON(http.StatusBadRequest, result.MethodNotAllowed.WithKeyPair("uri", ctx.Request.URL.RequestURI()))
		})
	}

	// 输出 json 声明中间件
	engine.Use(jsonResponseMiddleware)

	// 限流每 1 秒限制 1000 个
	engine.Use(rateLimitMiddleware(time.Second, 1000, 1))

	return &Server{
		engine: engine,
	}
}

// 路由对象
func (server *Server) router(router Router) {
	groupRouter := server.engine.Group(router.GroupName)

	if AuthMiddleware == nil || AdminMiddleware == nil {
		panic("middleware not implement")
		return
	}

	// 权限类型
	switch router.RoleType {
	case RolePublic:
	case RoleAuth:
		groupRouter.Use(AuthMiddleware) // 授权中间件
	case RoleAdmin:
		groupRouter.Use(AdminMiddleware) // 管理中间件
	}

	// 注册路由回调
	for _, routerPath := range router.RouterPaths {
		switch routerPath.HttpMethod {
		case http.MethodGet:
			groupRouter.GET(routerPath.RelativePath, exceptionHandler(routerPath.HandlerFunc))
		case http.MethodPost:
			groupRouter.POST(routerPath.RelativePath, exceptionHandler(routerPath.HandlerFunc))
		case http.MethodPut:
			groupRouter.PUT(routerPath.RelativePath, exceptionHandler(routerPath.HandlerFunc))
		case http.MethodDelete:
			groupRouter.DELETE(routerPath.RelativePath, exceptionHandler(routerPath.HandlerFunc))
		default:
			panic("http method not supported")
		}
	}
}

// 路由对象组
func (server *Server) RouterGroup(groupName string, routers []Router) {
	for _, router := range routers {
		router.GroupName = groupName + router.GroupName
		server.router(router)
	}
}

// 启动服务
func (server *Server) Run(port int) {
	err := server.engine.Run(":" + strconv.Itoa(port))
	if err != nil {
		panic(err.Error())
	}
}
