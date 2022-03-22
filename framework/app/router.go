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
	HttpMethod   string          // 方法类型 GET、POST、PUT、DELETE
	RelativePath string          // 路径
	HandlerFunc  gin.HandlerFunc // 回调
}

// 路由对象
type Router struct {
	GroupName   string          // 组名
	Middleware  gin.HandlerFunc // 中间件
	RouterPaths []RouterPath    // 路由路径
}
