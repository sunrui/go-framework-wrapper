/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/20 23:20:20
 */

package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// 获取当前用户
func getUser(ctx *gin.Context) {
	id := ctx.Param("id")
	fmt.Println(id)

	//app-gin.Response(ctx, result.Result.WithData(api-api-user))
}
