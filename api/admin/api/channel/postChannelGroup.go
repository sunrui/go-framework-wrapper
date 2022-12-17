/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-12-17 21:51:21
 */

package channel

import (
	"fmt"
	"framework/app/result"
	"github.com/gin-gonic/gin"
)

type postChannelGroupReq struct {
}

// 获取当前用户
func (controller Controller) postChannelGroup(ctx *gin.Context) *result.Result {
	id := ctx.Param("id")
	fmt.Println(id)

	return &result.Ok
}
