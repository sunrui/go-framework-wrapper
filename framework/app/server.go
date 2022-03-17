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

// 启动服务
func (app *Server) Run(port int) {
	err := app.engine.Run(":" + strconv.Itoa(port))
	if err != nil {
		panic(err.Error())
	}
}
