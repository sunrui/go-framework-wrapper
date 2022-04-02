package demo

import (
	"medium-server-go/framework/app"
	"net/http"
)

// GetRouter 获取路由对象
func GetRouter() app.Router {
	return app.Router{
		GroupName:  "/sms",
		Middleware: nil,
		RouterPaths: []app.RouterPath{
			{
				HttpMethod:   http.MethodPost,
				RelativePath: "/",
				HandlerFunc:  postSms,
			},
			{
				HttpMethod:   http.MethodGet,
				RelativePath: "/",
				HandlerFunc:  getSms,
			},
			{
				HttpMethod:   http.MethodGet,
				RelativePath: "/:phone",
				HandlerFunc:  getSmsOne,
			},
		},
	}
}

// GetRouters 获取注册路由
func GetRouters() []app.Router {
	return []app.Router{
		GetRouter(),
	}
}
