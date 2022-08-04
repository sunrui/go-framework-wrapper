/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-04-25 20:46:17
 */

package token

import (
	"errors"
	"framework/config"
	"framework/proto/result"
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

// 密钥
var secret = config.Jwt().Secret

// 过期时间（秒）
var maxAge = config.Jwt().MaxAge

// 令牌名称
const name = "token"

// 生成 jwt 字符串
func encode(payload Payload) (string, error) {
	claims := Token{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + int64(maxAge)*1000,
		},
		payload,
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenClaims.SignedString(secret)
}

// 验证 jwt 字符串
func decode(tokenString string) (*Token, error) {
	tokenClaims, err := jwt.ParseWithClaims(tokenString, &Token{}, func(token *jwt.Token) (any, error) {
		return secret, nil
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
		ctx.SetCookie(name, token, maxAge, "/", "", false, true)
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
	if tokenString, err = ctx.Cookie(name); err != nil {
		// 从 header 中获取令牌
		if tokenString = ctx.GetHeader(name); tokenString == "" {
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
	if token, err := Get(ctx); err != nil {
		// 没有令牌不刷新
		return
	} else {
		// 距离过期时间（毫秒）
		expired := token.ExpiresAt - time.Now().Unix()

		// 根据过期时间距离自动刷新
		if expired <= int64(config.Jwt().AutoRefresh)*1000 {
			Write(ctx, token.Payload)
		}
	}
}

// Remove 移除令牌
func Remove(ctx *gin.Context) {
	ctx.SetCookie(name, "", -1, "/", "", false, true)
}
