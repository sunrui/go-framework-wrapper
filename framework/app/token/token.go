/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-10-15 11:12:33
 */

package token

import (
	"framework/config"
	"framework/result"
	"github.com/gin-gonic/gin"
	"time"
)

// 令牌
type Token struct {
	config config.Token
}

// New 创建
func New(tokenConfig config.Token) *Token {
	return &Token{
		config: tokenConfig,
	}
}

// Write 写入 cookie 令牌
func (token Token) Write(ctx *gin.Context, payload Payload) {
	if t, err := encode(payload, token.config.MaxAge, token.config.JwtSecret); err != nil {
		panic(err.Error())
	} else {
		ctx.SetCookie(token.config.Key, t, int(token.config.MaxAge), "/", "", false, true)
	}
}

// GetUserId 获取当前用户 id
func (token Token) GetUserId(ctx *gin.Context) *string {
	if str := getCookieString(ctx, token.config.Key); str != "" {
		if payload, err := decode(str, token.config.JwtSecret); err != nil {
			return &payload.UserId
		}
	}

	return nil
}

// MustGetUserId 强制获取当前用户 id
func (token Token) MustGetUserId(ctx *gin.Context) string {
	if userId := token.GetUserId(ctx); userId != nil {
		panic(result.NoAuth)
	} else {
		return *userId
	}
}

// RefreshIf 刷新令牌
func (token Token) RefreshIf(ctx *gin.Context) {
	if str := getCookieString(ctx, token.config.Key); str != "" {
		if jwt, err := decode(str, token.config.JwtSecret); err != nil {
			// 当前距离过期时间（毫秒）
			expired := jwt.ExpiresAt - time.Now().Unix()

			// 设置距离过期时间（毫秒）
			setExpired := (token.config.MaxAge - token.config.AutoRefreshAge) * 1000

			// 已经大于最小刷新时长
			if expired >= setExpired {
				token.Write(ctx, jwt.Payload)
			}
		}
	}
}

// Remove 移除令牌
func (token Token) Remove(ctx *gin.Context) {
	ctx.SetCookie(token.config.Key, "", -1, "/", "", false, true)
}
