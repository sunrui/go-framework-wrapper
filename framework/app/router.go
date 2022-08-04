/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/03 07:50:03
 */

package app

import (
	"framework/proto/response"
	"framework/proto/result"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RouterFunc func(ctx *gin.Context) result.Result

// Router 路由路径
type Router struct {
	HttpMethod   string     // 方法类型 GET、POST、PUT、DELETE
	RelativePath string     // 路径
	RouterFunc   RouterFunc // 回调
}

// RouterGroup 路由对象
type RouterGroup struct {
	GroupName   string          // 组名
	Middleware  gin.HandlerFunc // 中间件
	RouterPaths []Router        // 路由路径
}

// gin 回调
func handlerFunc(routerFunc RouterFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		r := routerFunc(ctx)
		response.Reply(ctx, r)
	}
}

// 注册路由
func registerRouter(engine *gin.Engine, router RouterGroup) {
	routerGroup := engine.Group(router.GroupName)

	// 启用中间件
	if router.Middleware != nil {
		routerGroup.Use(router.Middleware)
	}

	// 注册路由回调
	for _, routerPath := range router.RouterPaths {
		switch routerPath.HttpMethod {
		case http.MethodGet:
			routerGroup.GET(routerPath.RelativePath, handlerFunc(routerPath.RouterFunc))
		case http.MethodPost:
			routerGroup.POST(routerPath.RelativePath, handlerFunc(routerPath.RouterFunc))
		case http.MethodPut:
			routerGroup.PUT(routerPath.RelativePath, handlerFunc(routerPath.RouterFunc))
		case http.MethodDelete:
			routerGroup.DELETE(routerPath.RelativePath, handlerFunc(routerPath.RouterFunc))
		default:
			panic("http method not supported")
		}
	}
}
