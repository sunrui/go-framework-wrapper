/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-04-07 11:01:46
 */

package api

import (
	"medium-server-go/framework/app"
	"medium-server-go/generate/api/user/template"
)

// GetRouters 获取注册路由
func GetRouters() []app.Router {
	return []app.Router{
		template.GetRouter(),
	}
}
