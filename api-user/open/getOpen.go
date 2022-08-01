/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/31 21:00:31
 */

package open

import (
	"framework/proto/result"
	"framework/proto/token"
	"github.com/gin-gonic/gin"
	"service/core/open"
)

// 获取指定用户下所有入驻
func getOpen(ctx *gin.Context) result.Result {
	// 获取当前用户 id
	userId := token.MustGetUserId(ctx)

	// 获取当前用户下的入驻
	opens := open.GetOpen(userId)

	return result.Ok.WithData(opens)
}
