/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-04-25 22:14:59
 */

package auth

import (
	"framework/proto/response"
	"framework/proto/result"
	"framework/proto/token"
	"github.com/gin-gonic/gin"
)

// 获取令牌
func getToken(ctx *gin.Context) {
	// 获取用户令牌
	t, err := token.GetToken(ctx)
	if err != nil {
		response.New(ctx).Result(result.NotFound)
		return
	}

	response.New(ctx).Data(t)
}