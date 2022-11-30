/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-10-15 11:12:33
 */

package token

import (
	"errors"
	"framework/config"
	"framework/result"
	"github.com/gin-gonic/gin"
	"strings"
)

// Token 令牌
type Token struct {
	Config  config.Token
	Storage Storage
}

// New 创建
func New(tokenConfig config.Token, storage Storage) *Token {
	return &Token{
		Config:  tokenConfig,
		Storage: storage,
	}
}

// Write 写入 cookie 令牌
func (token Token) Write(ctx *gin.Context, payload Payload) {
	if value, err := token.Storage.Set(payload, token.Config.MaxAge); err != nil {
		panic(err.Error())
	} else {
		ctx.SetCookie(token.Config.Key, value, int(token.Config.MaxAge), "/", "", false, true)
	}
}

// GetPayload 获取当前负荷
func (token Token) GetPayload(ctx *gin.Context) (payload *Payload, ttl int64, err error) {
	value := func(ctx *gin.Context, key string) (value string) {
		// 从 cookie 中获取令牌
		if value, err = ctx.Cookie(key); err != nil {
			// 从 header 中获取令牌
			if value = ctx.GetHeader(key); value == "" {
				// 从 Authorization 中获取令牌
				if value = ctx.GetHeader("Authorization"); value != "" {
					prefix := "Bearer "
					if strings.Index(value, prefix) == 0 {
						value = value[len(prefix):]
					}
				}
			}
		}

		return value
	}(ctx, token.Config.Key)

	if value != "" {
		payload, ttl, err = token.Storage.Get(value)

		// 判断是否需要刷新令牌
		if ttl <= token.Config.AutoRefreshAge {
			token.Write(ctx, *payload)
		}

		return
	} else {
		return nil, 0, errors.New("no value")
	}
}

// MustGetUserId 强制获取当前用户 id
func (token Token) MustGetUserId(ctx *gin.Context) string {
	if payload, _, err := token.GetPayload(ctx); err != nil {
		panic(result.NoAuth)
	} else {
		return payload.UserId
	}
}

// Remove 移除
func (token Token) Remove(ctx *gin.Context) {
	ctx.SetCookie(token.Config.Key, "", -1, "/", "", false, true)
}
