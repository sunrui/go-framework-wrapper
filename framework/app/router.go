/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/03 07:50:03
 */

package app

import (
	"github.com/gin-gonic/gin"
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
