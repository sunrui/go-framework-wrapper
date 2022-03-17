/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author:sunrui
 * Date: 2022/01/03 07:51:03
 */

package app_gin

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"medium-server-go/common/result"
	"net/http"
	"time"
)

// 授权中间件
var AuthMiddleware gin.HandlerFunc

// 管理中间件
var AdminMiddleware gin.HandlerFunc

// 输出 json 声明中间件
func jsonResponseMiddleware(ctx *gin.Context) {
	ctx.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	ctx.Next()
}

// 流量限制中间件
//
// @fillInterval 间隔单位
// @capacity 令牌桶容量
// @quantum 每隔多久
func rateLimitMiddleware(fillInterval time.Duration, capacity, quantum int64) gin.HandlerFunc {
	bucket := ratelimit.NewBucketWithQuantum(fillInterval, capacity, quantum)

	return func(ctx *gin.Context) {
		if bucket.TakeAvailable(1) < 1 {
			ctx.JSON(http.StatusBadRequest, result.RateLimit)
			return
		}

		ctx.Next()
	}
}
