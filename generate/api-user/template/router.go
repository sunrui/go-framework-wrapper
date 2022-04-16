/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022/04/16 15:19:16
 */

package template

import (
	"framework/app"
	"net/http"
)

// GetRouter 获取路由对象
func GetRouter() app.Router {
	return app.Router{
		GroupName:  "/api-user/template",
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
