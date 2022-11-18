/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-15 21:22:41
 */

package middleware

import (
	"config"
	"framework/result"
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"time"
)

// 令牌桶
var bucket *ratelimit.Bucket

// RateLimit 流量限制中间件
func RateLimit(ctx *gin.Context) *result.Result[any] {
	if bucket.TakeAvailable(1) < 1 {
		ctx.Abort()
		return &result.Result[any]{
			Code: result.RateLimit,
			Data: result.KeyValueData("uri", ctx.Request.URL.RequestURI()),
		}
	}

	ctx.Next()
	return nil
}

// 初始化
func init() {
	bucket = ratelimit.NewBucketWithQuantum(time.Second, // 间隔单位
		config.Inst().RateLimit.Capacity, // 令牌桶容量
		config.Inst().RateLimit.Quantum,  // 每隔多久
	)
}
