/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-04-07 16:12:19
 */

package template

import (
	"framework/app"
	"net/http"
)

// GetRouter 获取路由对象
func GetRouter() app.Router {
	return app.Router{
		GroupName:  "/admin/template",
		Middleware: nil,
		RouterPaths: []app.RouterPath{
			{
				HttpMethod:   http.MethodGet,
				RelativePath: "/:id",
				HandlerFunc:  getOne,
			}, {
				HttpMethod:   http.MethodGet,
				RelativePath: "/",
				HandlerFunc:  getAll,
			}, {
				HttpMethod:   http.MethodPut,
				RelativePath: "/",
				HandlerFunc:  putOne,
			}, {
				HttpMethod:   http.MethodDelete,
				RelativePath: "/:id",
				HandlerFunc:  deleteOne,
			},
		},
	}
}
