/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-26 13:37:48
 */

package user

import (
	"framework/server"
)

// GetRouter 获取路由
func GetRouter() server.RouterGroup {
	return server.RouterGroup{
		GroupName:  "/user",
		Middleware: nil,
		Routers:    []server.Router{},
	}
}
