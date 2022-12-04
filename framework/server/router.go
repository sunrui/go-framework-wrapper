/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-14 22:30:23
 */

package server

import (
	"framework/app/result"
	"github.com/gin-gonic/gin"
	"net/http"
)

// RouterFunc 路由回调
type RouterFunc func(ctx *gin.Context) *result.Result

// Router 路由路径
type Router struct {
	HttpMethod   string     // 方法类型 GET、POST、PUT、DELETE
	RelativePath string     // 路径
	RouterFunc   RouterFunc // 回调
}

// RouterGroup 路由
type RouterGroup struct {
	GroupName  string          // 组名
	Middleware gin.HandlerFunc // 中间件
	Routers    []Router        // 路由路径
}

// 路由回调
func (server Server) routerFunc(routerFunc RouterFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if r := routerFunc(ctx); r != nil {
			server.response(ctx, r)
		}
	}
}

// 注册路由
func (server Server) registerRouter(engine *gin.Engine, routerGroup RouterGroup) {
	group := engine.Group(routerGroup.GroupName)

	// 启用中间件
	if routerGroup.Middleware != nil {
		group.Use(routerGroup.Middleware)
	}

	// 注册路由回调
	for _, routerPath := range routerGroup.Routers {
		switch routerPath.HttpMethod {
		case http.MethodGet:
			group.GET(routerPath.RelativePath, server.routerFunc(routerPath.RouterFunc))
		case http.MethodPost:
			group.POST(routerPath.RelativePath, server.routerFunc(routerPath.RouterFunc))
		case http.MethodPut:
			group.PUT(routerPath.RelativePath, server.routerFunc(routerPath.RouterFunc))
		case http.MethodDelete:
			group.DELETE(routerPath.RelativePath, server.routerFunc(routerPath.RouterFunc))
		default:
			panic("http method not supported")
		}
	}
}
