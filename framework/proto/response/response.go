/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022/04/03 00:05:03
 */

package response

import (
	"framework/proto/result"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Reply 数据回复对象
func Reply(ctx *gin.Context, result result.Result) {
	// 返回客户端
	ctx.JSON(http.StatusOK, result)

	// 异步记录出错日志
	go func() {
		logResult(ctx, result)
	}()
}
