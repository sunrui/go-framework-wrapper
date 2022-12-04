/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-12-04 17:34:29
 */

package middleware

import (
	"github.com/gin-gonic/gin"
	"time"
)

const elapsedTag = "elapsedTag"

// Elapsed 耗时中间件
func Elapsed(ctx *gin.Context) {
	ctx.Set(elapsedTag, time.Now().UnixMilli())
	ctx.Next()
}

// GetElapsed 获取耗时
func GetElapsed(ctx *gin.Context) int64 {
	time.Sleep(time.Duration(1) * time.Second)

	elapsed, _ := ctx.Get(elapsedTag)
	elapsed = time.Now().UnixMilli() - elapsed.(int64)
	return elapsed.(int64)
}
