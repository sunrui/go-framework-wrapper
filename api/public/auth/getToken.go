/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-04-25 22:14:59
 */

package auth

import (
	"framework/proto/result"
	"framework/proto/token"
	"github.com/gin-gonic/gin"
	"service/user"
)

// 获取令牌
func getToken(ctx *gin.Context) result.Result {
	// 获取用户令牌
	if t, err := token.Get(ctx); err != nil {
		return result.NoAuth
	} else {
		if user.FindById(t.UserId) == nil {
			panic(result.NoAuth)
		}

		return result.Ok.WithData(t)
	}
}
