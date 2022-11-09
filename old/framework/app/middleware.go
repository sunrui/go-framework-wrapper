/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/01
 */

package app

import (
	"framework/app/middleware"
	"github.com/gin-gonic/gin"
)

// 注册中间件
func registerMiddleware(engine *gin.Engine) {
	// 注册 404 回调
	engine.NoRoute(middleware.NotFound)

	// 注册 405 回调
	engine.HandleMethodNotAllowed = true
	engine.NoMethod(middleware.MethodNotAllowed)

	// 注册限流中间件
	engine.Use(middleware.RateLimit)

	// 注册文档中间件
	engine.GET("/doc/*any", middleware.Swagger)

	// 注册刷新令牌中间件
	engine.Use(middleware.Token)

	// 注册异常中间件
	engine.Use(middleware.Recover)

	// 注册 body 中间件
	engine.Use(middleware.Body)
}
