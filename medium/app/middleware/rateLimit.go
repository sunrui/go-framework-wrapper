/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-04-25 22:13:05
 */

package middleware

import (
	"github.com/gin-gonic/gin"
)

// RateLimit 流量限制中间件
func RateLimit(ctx *gin.Context) {
	//bucket := ratelimit.NewBucketWithQuantum(time.Second, // 间隔单位
	//	config.Cur().RateLimit.Capacity, // 令牌桶容量
	//	config.Cur().RateLimit.Quantum,  // 每隔多久
	//)

	//if bucket.TakeAvailable(1) < 1 {
	//ctx.JSON(http.StatusOK, result.RateLimit)
	//ctx.Abort()
	//return
	//}

	ctx.Next()
}
