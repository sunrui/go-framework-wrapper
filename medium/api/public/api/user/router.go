/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-26 13:37:48
 */

package user

import (
	"framework/http"
)

// GetRouter 获取路由
func GetRouter() http.RouterGroup {
	return http.RouterGroup{
		GroupName:  "/user",
		Middleware: nil,
		Routers:    []http.Router{},
	}
}
