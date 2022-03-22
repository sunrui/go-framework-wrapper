/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/25 22:29:25
 */

package public

import (
	"medium-server-go/api/public/area"
	"medium-server-go/api/public/auth"
	"medium-server-go/api/public/sms"
	"medium-server-go/framework/app"
)

// 获取注册路由
func GetRouters() []app.Router {
	return []app.Router{
		area.GetRouter(),
		sms.GetRouter(),
		auth.GetRouter(),
	}
}
