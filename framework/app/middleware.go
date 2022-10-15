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
	engine.NoRoute(middleware.NotFoundMiddleware)

	// 注册 405 回调
	engine.HandleMethodNotAllowed = true
	engine.NoMethod(middleware.MethodNotAllowedMiddleware)

	// 注册限流中间件
	engine.Use(middleware.RateLimitMiddleware)

	// 注册文档中间件
	engine.GET("/doc/*any", middleware.DocMiddleware)

	// 注册刷新令牌中间件
	engine.Use(middleware.TokenMiddleware)

	// 注册异常中间件
	engine.Use(middleware.RecoverMiddleware)

	// 注册 body 中间件
	engine.Use(middleware.BodyMiddleware)
}
