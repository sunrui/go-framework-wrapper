/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022/04/03 00:05:03
 */

package response

import (
	"framework/config"
	"framework/log"
	"framework/proto/request"
	"framework/proto/result"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Reply 回复
func Reply(ctx *gin.Context, result result.Result) {
	// 是否需要导出请求
	if request.IsExport(ctx) {
		req := request.GetRequest(ctx)
		result.Request = &req
	}

	// 返回客户端
	ctx.JSON(http.StatusOK, result)

	// 异步记录日志
	if config.Log().Enable {
		go func() {
			log.WriteResult(ctx, result)
		}()
	}
}
