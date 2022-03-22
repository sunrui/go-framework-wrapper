package middleware

import (
	"github.com/gin-gonic/gin"
	"medium-server-go/framework/app"
	"medium-server-go/framework/result"
	"medium-server-go/framework/token"
)

// 授权中间件
func AuthMiddleware(ctx *gin.Context) {
	_, err := token.GetTokenEntity(ctx)
	if err != nil {
		app.Response(ctx, result.NoAuth)
	}
}
