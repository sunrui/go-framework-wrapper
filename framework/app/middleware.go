package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"io/ioutil"
	"medium-server-go/framework/config"
	"medium-server-go/framework/exception"
	"os"
	"path/filepath"
	"runtime"
	"runtime/debug"
	"strings"
	"time"
)

// 异常 404 中间件
func notFoundMiddleware(ctx *gin.Context) {
	Result(ctx).Exception(exception.NotFound.WithKeyPair("uri", ctx.Request.URL.RequestURI()))
}

// 异常 405 中间件
func methodNotAllowed(ctx *gin.Context) {
	Result(ctx).Exception(exception.MethodNotAllowed.WithKeyPair("uri", ctx.Request.URL.RequestURI()))
}

// 流量限制中间件
//
// @fillInterval  间隔单位
// @capacity      令牌桶容量
// @quantum       每隔多久
func ratelimitMiddleware(fillInterval time.Duration, capacity, quantum int64) gin.HandlerFunc {
	bucket := ratelimit.NewBucketWithQuantum(fillInterval, capacity, quantum)

	return func(ctx *gin.Context) {
		if bucket.TakeAvailable(1) < 1 {
			Result(ctx).Exception(exception.RateLimit)
			return
		}

		ctx.Next()
	}
}

// swagger 文档中间件
func redocMiddleware(ctx *gin.Context) {
	suffix := filepath.Base(ctx.Request.URL.Path)
	switch suffix {
	case "doc.json":
		data, _ := ioutil.ReadFile("docs/swagger.json")
		_, _ = ctx.Writer.Write(data)
		return
	case "index.html":
		_, _ = ctx.Writer.Write([]byte(`
					<!DOCTYPE html>
					<html>
					  <head>
						<meta charset="utf-8"/>
						<meta name="viewport" content="width=device-width, initial-scale=1">
						<style>
						  body {
							margin: 0;
							padding: 0;
						  }
						</style>
					  </head>
					  <body>
						<redoc spec-url='swagger/doc.json'></redoc>
						<script src="https://cdn.jsdelivr.net/npm/redoc@latest/bundles/redoc.standalone.js"> </script>
					  </body>
					</html>
					`))
		return
	default:
		Result(ctx).Exception(exception.NotFound)
	}
}

// 堆栈对象
type stack struct {
	Function string `json:"function"` // 函数
	File     string `json:"file"`     // 行数
}

// 获取推栈层级
func getStack() stack {
	// 最大函数层级 5
	pc := make([]uintptr, 5)
	runtime.Callers(3, pc)
	frames := runtime.CallersFrames(pc)

	// 当前项目目录
	pwd, _ := os.Getwd()

	for frame, ok := frames.Next(); ok; frame, ok = frames.Next() {
		// 过滤掉系统目录
		if !strings.HasPrefix(frame.File, pwd) {
			continue
		}

		// 去掉项目目录
		file := strings.Replace(frame.File, pwd, "", -1)
		file = fmt.Sprintf("%s:%d", file, frame.Line)
		function := filepath.Base(frame.Function)

		return stack{
			Function: function,
			File:     file,
		}
	}

	return stack{}
}

// 异常捕获中间件
func recoverMiddleware(ctx *gin.Context) {
	// 捕获对象，全部抛出可以使用 panic 方法。
	defer func() {
		if err := recover(); err != nil {
			dataMap := make(map[string]interface{})
			dataMap["stack"] = getStack()

			// 判断是否抛出了 exception 对象
			res, ok := err.(*exception.Exception)
			if ok {
				dataMap["error"] = res.Data
			} else {
				dataMap["error"] = err
			}

			Result(ctx).Exception(exception.InternalError.WithData(dataMap))

			// 为了更好的调试，在开发环境中输出系统错误。
			if config.IsDebugMode() {
				debug.PrintStack()
			}
		}
	}()

	ctx.Next()
}

// body 中间件
func bodyMiddleware(ctx *gin.Context) {
	ctx.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	ctx.Next()
}

// 注册中间件
func registerMiddleware(engine *gin.Engine) {
	// 注册 404 回调
	engine.NoRoute(notFoundMiddleware)

	// 注册 405 回调
	engine.HandleMethodNotAllowed = true
	engine.NoMethod(methodNotAllowed)

	// 注册限流中间件
	engine.Use(ratelimitMiddleware(time.Second, 200, 1))

	// 注册文档中间件
	engine.GET("/swagger/*any", redocMiddleware)

	// 注册异常中间件
	engine.Use(recoverMiddleware)

	// 注册 body 中间件
	engine.Use(bodyMiddleware)
}
