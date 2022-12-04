/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-12-05 04:10:16
 */

package log

import (
	"framework/app/mysql"
	"framework/server"
	"medium/service/log"
	"net/http"
)

type Controller struct {
	Mysql          *mysql.Mysql
	HttpRepository log.HttpRepository
}

func NewController(mysql *mysql.Mysql) Controller {
	return Controller{
		Mysql:          mysql,
		HttpRepository: log.NewHttpRepository(mysql),
	}
}

// GetRouter 获取路由
func (controller Controller) GetRouter() server.RouterGroup {
	return server.RouterGroup{
		GroupName:  "/log",
		Middleware: nil,
		Routers: []server.Router{
			{
				HttpMethod:   http.MethodGet,
				RelativePath: "/",
				RouterFunc:   controller.getIndex,
			},
		},
	}
}
