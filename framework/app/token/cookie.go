/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-26 19:43:36
 */

package token

import (
	"github.com/gin-gonic/gin"
	"strings"
)

// 获取当前用户令牌
func getCookieString(ctx *gin.Context, key string) string {
	var str string
	var err error

	// 从 cookie 中获取令牌
	if str, err = ctx.Cookie(key); err != nil {
		// 从 header 中获取令牌
		if str = ctx.GetHeader(key); str == "" {
			// 从 Authorization 中获取令牌
			if str = ctx.GetHeader("Authorization"); str != "" {
				prefix := "Bearer "
				if strings.Index(str, prefix) == 0 {
					str = str[len(prefix):]
				}
			}
		}
	}

	return str
}
