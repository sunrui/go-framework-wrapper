/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-15 21:22:41
 */

package middleware

import (
	"framework/app/result"
	"github.com/didip/tollbooth/v7"
	"github.com/didip/tollbooth/v7/limiter"
	"github.com/gin-gonic/gin"
	"time"
)

// RateLimit 限流
type RateLimit struct {
	lmt *limiter.Limiter
}

// NewRateLimit 创建限流
func NewRateLimit(capacity float64) RateLimit {
	lmt := tollbooth.NewLimiter(capacity, &limiter.ExpirableOptions{DefaultExpirationTTL: 24 * time.Hour})
	lmt.SetIPLookups([]string{"X-Forwarded-For", "X-Real-IP", "RemoteAddr"})

	return RateLimit{
		lmt: lmt,
	}
}

// Take 监测限流
func (rateLimit RateLimit) Take(ctx *gin.Context) {
	if err := tollbooth.LimitByRequest(rateLimit.lmt, ctx.Writer, ctx.Request); err != nil {
		ctx.Abort()

		panic(result.RateLimit.WithData(result.M{
			"uri": ctx.Request.URL.RequestURI(),
		}))
	} else {
		ctx.Next()
	}
}
