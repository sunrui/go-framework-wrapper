/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/01
 */

package app

import (
	"framework/app/log"
	"framework/app/middleware"
	"framework/result"
	"github.com/gin-gonic/gin"
	"net/http"
)

// MiddlewareFunc 中间件回调对象
type MiddlewareFunc func(ctx *gin.Context) *result.Result[any]

// 中间件回调
func middlewareFunc(middlewareFunc MiddlewareFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		r := middlewareFunc(ctx)
		if r != nil {
			ctx.JSON(http.StatusOK, r)

			defer func() {
				log.Write(ctx, *r)
			}()
		}
	}
}

// 注册中间件
func registerMiddleware(engine *gin.Engine) {
	// 注册 404 回调
	engine.NoRoute(middlewareFunc(middleware.NotFound))

	// 注册 405 回调
	engine.HandleMethodNotAllowed = true
	engine.NoMethod(middlewareFunc(middleware.MethodNotAllowed))

	// 注册限流中间件
	engine.Use(middlewareFunc(middleware.RateLimit))

	// 注册异常中间件
	engine.Use(middlewareFunc(middleware.Recover))

	// 注册刷新令牌中间件
	engine.Use(middleware.Token)

	// 注册文档中间件
	engine.GET("/doc/*any", middleware.Swagger)

	// 注册 body 中间件
	engine.Use(middleware.Body)
}
