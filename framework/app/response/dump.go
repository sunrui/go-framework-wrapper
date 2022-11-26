/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-26 20:39:54
 */

package response

import (
	"framework/context"
	"github.com/gin-gonic/gin"
)

// IsDump 是否导出
func IsDump(ctx *gin.Context) bool {
	dump := ctx.DefaultQuery("dump", "")
	if dump == "false" || dump == "0" {
		return false
	}

	return context.Config.Request.Dump
}

// SetDump 设置导出
func SetDump(dump bool) {
	context.Config.Request.Dump = dump
}
