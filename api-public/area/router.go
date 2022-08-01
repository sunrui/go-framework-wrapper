/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/31 14:44:31
 */

package area

import (
	"framework/app"
	"net/http"
)

// GetRouter 获取路由对象
func GetRouter() app.RouterGroup {
	return app.RouterGroup{
		GroupName:  "/area",
		Middleware: nil,
		RouterPaths: []app.Router{
			{
				HttpMethod:   http.MethodGet,
				RelativePath: "",
				RouterFunc:   getCountry,
			},
			{
				HttpMethod:   http.MethodGet,
				RelativePath: "/province",
				RouterFunc:   getProvince,
			},
			{
				HttpMethod:   http.MethodGet,
				RelativePath: "/province/:provinceId",
				RouterFunc:   getCity,
			},
			{
				HttpMethod:   http.MethodGet,
				RelativePath: "/city/:cityId",
				RouterFunc:   getArea,
			},
		},
	}
}
