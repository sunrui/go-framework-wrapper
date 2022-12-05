/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-15 21:22:41
 */

package middleware

import (
	"framework/app/result"
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"time"
)

// RateLimit 限流
type RateLimit struct {
	*ratelimit.Bucket
}

// NewRateLimit 创建限流
func NewRateLimit(capacity int64, quantum int64) RateLimit {
	return RateLimit{
		Bucket: ratelimit.NewBucketWithQuantum(time.Second, // 间隔单位
			capacity, // 令牌桶容量
			quantum,  // 每隔多久
		),
	}
}

// Take 监测限流
func (bucket RateLimit) Take(ctx *gin.Context) *result.Result {
	if bucket.TakeAvailable(1) < 1 {
		ctx.Abort()

		return result.RateLimit.WithData(result.M{
			"uri": ctx.Request.URL.RequestURI(),
		})
	}

	ctx.Next()
	return nil
}
