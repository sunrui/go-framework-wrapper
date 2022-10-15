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
	"github.com/golang-jwt/jwt"
	"strings"
	"time"
)

// Payload 对象
type Payload struct {
	UserId string `json:"userId"` // 用户 id
}

// Token 对象
type Token struct {
	jwt.StandardClaims
	Payload
}

// 生成 jwt 字符串
func encode(payload Payload) (string, error) {
	token := Token{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + int64(config.Get().Jwt.MaxAge)*1000,
		},
		payload,
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, token)
	return tokenClaims.SignedString(config.Get().Jwt.Secret)
}

// 验证 jwt 字符串
func decode(tokenString string) (*Token, error) {
	tokenClaims, err := jwt.ParseWithClaims(tokenString, &Token{}, func(token *jwt.Token) (any, error) {
		return config.Get().Jwt.Secret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Token); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

// Write 写入 cookie 令牌
func Write(ctx *gin.Context, payload Payload) {
	if token, err := encode(payload); err != nil {
		panic(err.Error())
	} else {
		ctx.SetCookie(config.Get().Jwt.Key, token, config.Get().Jwt.MaxAge, "/", "", false, true)
	}
}

// MustGetUserId 获取当前用户 id
func MustGetUserId(ctx *gin.Context) string {
	if token, err := Get(ctx); err != nil {
		panic(result.NoAuth)
	} else {
		return token.UserId
	}
}

// Get 获取当前用户令牌
func Get(ctx *gin.Context) (*Token, error) {
	var tokenString string
	var err error

	// 从 cookie 中获取令牌
	if tokenString, err = ctx.Cookie(config.Get().Jwt.Key); err != nil {
		// 从 header 中获取令牌
		if tokenString = ctx.GetHeader(config.Get().Jwt.Key); tokenString == "" {
			// 从 Authorization 中获取令牌
			if tokenString = ctx.GetHeader("Authorization"); tokenString != "" {
				prefix := "Bearer "
				if strings.Index(tokenString, prefix) == 0 {
					tokenString = tokenString[len(prefix):]
				}
			}
		}
	}

	if tokenString == "" {
		return nil, errors.New("<null>")
	}

	return decode(tokenString)
}

// RefreshIf 刷新令牌
func RefreshIf(ctx *gin.Context) {
	if token, err := Get(ctx); err == nil {
		// 距离过期时间（毫秒）
		expired := token.ExpiresAt - time.Now().Unix()

		// 根据过期时间距离自动刷新
		if expired <= int64(config.Get().Jwt.MaxAge)*100-int64(config.Get().Jwt.AutoRefresh)*1000 {
			Write(ctx, token.Payload)
		}
	}
}

// Remove 移除令牌
func Remove(ctx *gin.Context) {
	ctx.SetCookie(config.Get().Jwt.Key, "", -1, "/", "", false, true)
}
