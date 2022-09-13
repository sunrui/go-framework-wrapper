/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-04-25 22:13:05
 */

package app

import (
	"fmt"
	"framework/config"
	"framework/proto/doc"
	"framework/proto/request"
	"framework/proto/response"
	"framework/proto/result"
	"framework/proto/token"
	"framework/utils"
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"net/http"
	"path/filepath"
	"strings"
	"time"
)

// 异常 404 中间件
func notFoundMiddleware(ctx *gin.Context) {
	response.Reply(ctx, result.NotFound.WithKeyPair("uri", ctx.Request.URL.RequestURI()))
	ctx.Abort()
}

// 异常 405 中间件
func methodNotAllowedMiddleware(ctx *gin.Context) {
	response.Reply(ctx, result.MethodNotAllowed.WithKeyPair("uri", ctx.Request.URL.RequestURI()))
	ctx.Abort()
}

// 流量限制中间件
func rateLimitMiddleware(ctx *gin.Context) {
	bucket := ratelimit.NewBucketWithQuantum(time.Second, // 间隔单位
		config.Get().RateLimit.Capacity, // 令牌桶容量
		config.Get().RateLimit.Quantum,  // 每隔多久
	)

	if bucket.TakeAvailable(1) < 1 {
		response.Reply(ctx, result.RateLimit)
		ctx.Abort()
		return
	}

	ctx.Next()
}

// swagger 文档中间件
func docMiddleware(ctx *gin.Context) {
	path := ctx.Request.URL.Path

	// 非 /doc 开头不是文档
	if !strings.HasPrefix(path, "/doc/") {
		return
	}

	// 过滤掉非法的 /doc/? 路径
	suffix := filepath.Base(path)
	if suffix != "doc" && suffix != "doc.json" && suffix != "redoc.js" {
		ctx.Redirect(http.StatusFound, "/doc")
		return
	}

	_, _ = ctx.Writer.Write(doc.Redoc(suffix))
}

// 注册刷新令牌中间件
func tokenMiddleware(ctx *gin.Context) {
	token.RefreshIf(ctx)
	ctx.Next()
}

// 异常捕获中间件
func recoverMiddleware(ctx *gin.Context) {
	// 捕获对象，全部抛出可以使用 panic 方法。
	defer func() {
		if err := recover(); err != nil {
			// 判断是否抛出了 Result 对象
			if res, ok := err.(result.Result); ok {
				response.Reply(ctx, res)
			} else {
				// 堆栈信息
				dataMap := make(map[string]any)
				dataMap["stack"] = utils.DumpStack(5)
				dataMap["error"] = fmt.Sprintf("%s", err)

				r := result.InternalError.WithData(dataMap)
				response.Reply(ctx, r)
				ctx.Abort()
			}
		}
	}()

	ctx.Next()
}

// body 中间件
func bodyMiddleware(ctx *gin.Context) {
	// 如果需要记录日志或请求被异出则拷贝 body 对象
	if config.Get().Log.Enable || request.IsRequestDump(ctx) {
		request.CopyBody(ctx)
	}

	ctx.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	ctx.Next()
}

// 注册中间件
func registerMiddleware(engine *gin.Engine) {
	// 注册 404 回调
	engine.NoRoute(notFoundMiddleware)

	// 注册 405 回调
	engine.HandleMethodNotAllowed = true
	engine.NoMethod(methodNotAllowedMiddleware)

	// 注册限流中间件
	engine.Use(rateLimitMiddleware)

	// 注册文档中间件
	engine.GET("/doc/*any", docMiddleware)

	// 注册刷新令牌中间件
	engine.Use(tokenMiddleware)

	// 注册异常中间件
	engine.Use(recoverMiddleware)

	// 注册 body 中间件
	engine.Use(bodyMiddleware)
}
