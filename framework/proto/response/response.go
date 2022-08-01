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

// Reply 数据返回对象
func Reply(ctx *gin.Context, result result.Result) {
	ctx.JSON(http.StatusOK, result)
}
