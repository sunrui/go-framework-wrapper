/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/03 07:50:03
 */

package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// RouterPath 路由路径
type RouterPath struct {
	HttpMethod   string          // 方法类型 GET、POST、PUT、DELETE
	RelativePath string          // 路径
	HandlerFunc  gin.HandlerFunc // 回调
}

// Router 路由对象
type Router struct {
	GroupName   string          // 组名
	Middleware  gin.HandlerFunc // 中间件
	RouterPaths []RouterPath    // 路由路径
}

// Router 路由对象
func registerRouter(engine *gin.Engine, router Router) {
	groupRouter := engine.Group(router.GroupName)

	// 启用中间件
	if router.Middleware != nil {
		groupRouter.Use(router.Middleware)
	}

	// 注册路由回调
	for _, routerPath := range router.RouterPaths {
		switch routerPath.HttpMethod {
		case http.MethodGet:
			groupRouter.GET(routerPath.RelativePath, routerPath.HandlerFunc)
		case http.MethodPost:
			groupRouter.POST(routerPath.RelativePath, routerPath.HandlerFunc)
		case http.MethodPut:
			groupRouter.PUT(routerPath.RelativePath, routerPath.HandlerFunc)
		case http.MethodDelete:
			groupRouter.DELETE(routerPath.RelativePath, routerPath.HandlerFunc)
		default:
			panic("http method not supported")
		}
	}
}
