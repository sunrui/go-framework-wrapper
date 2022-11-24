/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/20 23:20:20
 */

package user

import (
	"fmt"
	"framework/result"
	"github.com/gin-gonic/gin"
)

// 获取当前用户
func getUser(ctx *gin.Context) result.Result[any] {
	id := ctx.Param("id")
	fmt.Println(id)

	return result.Result[any]{}
}
