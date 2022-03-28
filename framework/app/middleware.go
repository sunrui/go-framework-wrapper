/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author:sunrui
 * Date: 2022/01/03 07:51:03
 */

package app

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"io/ioutil"
	"medium-server-go/framework/exception"
	"path/filepath"
	"time"
)

// 输出 json 声明中间件
func jsonResponseMiddleware(ctx *gin.Context) {
	ctx.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	ctx.Next()
}

// 流量限制中间件
//
// @fillInterval  间隔单位
// @capacity      令牌桶容量
// @quantum       每隔多久
func rateLimitMiddleware(fillInterval time.Duration, capacity, quantum int64) gin.HandlerFunc {
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
