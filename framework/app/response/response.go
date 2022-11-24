/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-24 22:35:55
 */

package response

import (
	"fmt"
	"framework/app/log"
	"framework/app/request"
	"framework/config"
	"framework/result"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response 返回
func Response(ctx *gin.Context, r *result.Result) {
	// 结果导出请求
	if request.IsDump(ctx) {
		req := request.Get(ctx)
		r.Request = &req
	}

	go func() {
		// 开启控制台
		if config.IsDev() {
			fmt.Println(r)
		}

		// 记录日志
		log.WriteResult(ctx, r)
	}()

	// 返回客户端
	ctx.AbortWithStatusJSON(http.StatusOK, r)
}
