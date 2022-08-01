/*
 * Copyright (c) $today.year honeysense.com All rights reserved.
 * Author: $author
 * Date: $today.format("yyyy-MM-dd HH:mm:ss")
 */

package template

import (
	"framework/app"
	"net/http"
)

// GetRouter 获取路由对象
func GetRouter() app.RouterGroup {
	return app.RouterGroup{
		GroupName:  "/api-user/template",
		Middleware: nil,
		RouterPaths: []app.Router{
			{
				HttpMethod:   http.MethodGet,
				RelativePath: "/:id",
				RouterFunc:   getOne,
			}, {
				HttpMethod:   http.MethodGet,
				RelativePath: "/",
				RouterFunc:   getAll,
			}, {
				HttpMethod:   http.MethodPost,
				RelativePath: "/",
				RouterFunc:   postOne,
			}, {
				HttpMethod:   http.MethodPut,
				RelativePath: "/:id",
				RouterFunc:   putOne,
			}, {
				HttpMethod:   http.MethodDelete,
				RelativePath: "/:id",
				RouterFunc:   deleteOne,
			},
		},
	}
}
