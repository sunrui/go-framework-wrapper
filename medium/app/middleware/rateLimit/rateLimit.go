/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-15 01:11:11
 */

package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"medium/result"
	"time"
)

// 令牌桶容量
var capacity int64 = 1000

// 每隔多少秒
var quantum int64 = 1

// RateLimit 流量限制中间件
func RateLimit(ctx *gin.Context) *result.Result[any] {
	bucket := ratelimit.NewBucketWithQuantum(time.Second, // 间隔单位
		capacity, // 令牌桶容量
		quantum,  // 每隔多久
	)

	if bucket.TakeAvailable(1) < 1 {
		ctx.Abort()
		return &result.Result[any]{
			Code: result.RATE_LIMIT,
			Data: result.KeyValueData("uri", ctx.Request.URL.RequestURI()),
		}
	}

	ctx.Next()
	return nil
}

// SetRateLimit 设置限流
func SetRateLimit(_capacity int64, _quantum int64) {
	capacity = _capacity
	quantum = _quantum
}
