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

// 路由路径
type RouterPath struct {
	// 方法类型 GET、POST、PUT、DELETE
	HttpMethod string
	// 路径
	RelativePath string
	// 回调
	HandlerFunc gin.HandlerFunc
}

// 权限类型
type RoleType int

const (
	RolePublic  = iota // 开放权限
	RoleAuth           // 登录权限
	RoleAdmin          // 管理权限
	RoleChannel        // 渠道权限
)

// 路由对象
type Router struct {
	// 组名
	GroupName string
	// 权限类型
	RoleType RoleType
	// 路由路径
	RouterPaths []RouterPath
}

// 注册路由对象
func (app *Server) RegisterRouter(router Router) {
	groupRouter := app.engine.Group(router.GroupName)

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
