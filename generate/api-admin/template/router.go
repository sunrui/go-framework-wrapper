/*
 * Copyright (c) $today.year honeysense.com All rights reserved.
 * Author: sunrui
 * Date: $today.format("yyyy-MM-dd HH:mm:ss")
 */

package template

import (
	"framework/app"
	"net/http"
)

// GetRouter 获取路由对象
func GetRouter() app.Router {
	return app.Router{
		GroupName:  "/api-admin/template",
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
				HttpMethod:   http.MethodPost,
				RelativePath: "/",
				HandlerFunc:  postOne,
			}, {
				HttpMethod:   http.MethodPut,
				RelativePath: "/:id",
				HandlerFunc:  putOne,
			}, {
				HttpMethod:   http.MethodDelete,
				RelativePath: "/:id",
				HandlerFunc:  deleteOne,
			},
		},
	}
}
