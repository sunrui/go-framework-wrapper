/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-12-17 20:42:11
 */

package channel

import (
	"fmt"
	"framework/app/result"
	"github.com/gin-gonic/gin"
)

type postChannelReq struct {
}

// 获取当前用户
func (controller Controller) postChannel(ctx *gin.Context) *result.Result {
	id := ctx.Param("id")
	fmt.Println(id)

	return &result.Ok
}
