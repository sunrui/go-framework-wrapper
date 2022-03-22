/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/31 14:44:31
 */

package area

import (
	"medium-server-go/framework/app"
	"net/http"
)

// 获取路由对象
func GetRouter() app.Router {
	return app.Router{
		GroupName:  "/area",
		Middleware: nil,
		RouterPaths: []app.RouterPath{
			{
				HttpMethod:   http.MethodGet,
				RelativePath: "",
				HandlerFunc:  getCountry,
			},
			{
				HttpMethod:   http.MethodGet,
				RelativePath: "/province",
				HandlerFunc:  getProvince,
			},
			{
				HttpMethod:   http.MethodGet,
				RelativePath: "/province/:provinceId",
				HandlerFunc:  getCity,
			},
			{
				HttpMethod:   http.MethodGet,
				RelativePath: "/city/:cityId",
				HandlerFunc:  getArea,
			},
		},
	}
}
