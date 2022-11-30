/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-24 22:35:55
 */

package response

import (
	"fmt"
	"framework/config"
	"framework/context"
	"framework/http/request"
	"framework/result"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response 返回
func Response(ctx *gin.Context, r *result.Result) {
	// 结果导出请求
	if IsDump(ctx) {
		req := request.Get(ctx)
		r.Request = &req
	}

	go func() {
		// 控制台日志
		if config.IsDev() {
			fmt.Println(r)
		}

		// 记录日志
		buffer := getResultString(ctx, r)
		if r.Code == result.Ok.Code {
			if context.Log.HttpAccess != nil {
				context.Log.HttpAccess.Debugln(buffer)
			}
		} else {
			if context.Log.HttpError != nil {
				context.Log.HttpError.Errorln(buffer)
			}
		}
	}()

	// 返回客户端
	ctx.AbortWithStatusJSON(http.StatusOK, r)
}
