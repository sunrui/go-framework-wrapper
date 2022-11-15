/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-14 22:30:23
 */

package app

import (
	"framework/app/log"
	"framework/app/request"
	"framework/result"
	"github.com/gin-gonic/gin"
	"net/http"
)

// RouterFunc 路由回调
type RouterFunc func(ctx *gin.Context) result.Result[any]

// Router 路由路径
type Router struct {
	HttpMethod   string     // 方法类型 GET、POST、PUT、DELETE
	RelativePath string     // 路径
	RouterFunc   RouterFunc // 回调
}

// RouterGroup 路由对象
type RouterGroup struct {
	GroupName  string          // 组名
	Middleware gin.HandlerFunc // 中间件
	Routers    []Router        // 路由路径
}

// 路由回调
func routerFunc(routerFunc RouterFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		r := routerFunc(ctx)

		// 结果导出请求
		if request.IsEnable() {
			req := request.Get(ctx)
			r.Request = &req
		}

		// 记录日志
		log.Write(ctx, r)

		// 返回客户端
		ctx.JSON(http.StatusOK, r)
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
	for _, routerPath := range router.Routers {
		switch routerPath.HttpMethod {
		case http.MethodGet:
			routerGroup.GET(routerPath.RelativePath, routerFunc(routerPath.RouterFunc))
		case http.MethodPost:
			routerGroup.POST(routerPath.RelativePath, routerFunc(routerPath.RouterFunc))
		case http.MethodPut:
			routerGroup.PUT(routerPath.RelativePath, routerFunc(routerPath.RouterFunc))
		case http.MethodDelete:
			routerGroup.DELETE(routerPath.RelativePath, routerFunc(routerPath.RouterFunc))
		default:
			panic("http method not supported")
		}
	}
}
