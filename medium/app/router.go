/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-14 22:30:23
 */

package app

import (
	"github.com/gin-gonic/gin"
	"medium/result"
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

// gin 回调
func handlerFunc(routerFunc RouterFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		r := routerFunc(ctx)

		//// 是否结果导出请求
		//if request.IsDebug(ctx) {
		//	req := request.Get(ctx)
		//	r.Request = &req
		//}
		//
		//// 记录日志
		//if config.Cur().Log.Enable {
		//	// 写文件
		//	if config.Cur().Log.WriteFile {
		//		if buffer := getBuffer(ctx, r); buffer != "" {
		//			log.Println(buffer)
		//		}
		//	}
		//
		//	// 写控制台
		//	if config.Cur().Log.WriteConsole {
		//		if buffer := getBuffer(ctx, r); buffer != "" {
		//			println(buffer)
		//		}
		//	}
		//}

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
