/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/25 22:29:25
 */

package user

import (
	"medium-server-go/api/user/open"
	"medium-server-go/api/user/user"
	"medium-server-go/framework/app"
)

// GetRouters 获取注册路由
func GetRouters() []app.Router {
	return []app.Router{
		open.GetRouter(),
		user.GetRouter(),
	}
}
