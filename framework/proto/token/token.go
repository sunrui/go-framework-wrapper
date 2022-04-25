/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-04-25 20:46:17
 */

package token

import (
	"framework/config"
	"framework/proto/result"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"strings"
)

// Payload 令牌对象
type Payload struct {
	UserId string `json:"userId"`
}

// jwt 对象
type jwtClaims struct {
	jwt.StandardClaims
	Payload
}

// jwt 密钥
var jwtSecret = config.Get().Jwt.Secret

// 令牌 key 名称
const tokenKey = "token"

// 令牌过期时间默认 30 天
const tokenMaxAge = 30 * 24 * 60 * 60

// 生成 jwt 字符串
func encode(payload Payload) (token string, err error) {
	claims := jwtClaims{
		jwt.StandardClaims{},
		payload,
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenClaims.SignedString(jwtSecret)
}

// 验证 jwt 字符串
func decode(token string) (payload *Payload, err error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &jwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*jwtClaims); ok && tokenClaims.Valid {
			return &claims.Payload, nil
		}
	}

	return nil, err
}

// Write 写入 cookie 令牌
func Write(ctx *gin.Context, userId string) {
	// 生成用户令牌
	token, err := encode(Payload{
		UserId: userId,
	})
	if err != nil {
		return
	}

	// 写入令牌，默认 30 天
	ctx.SetCookie(tokenKey, token, tokenMaxAge, "/", "", false, true)
}

// GetUserId 获取当前用户 id
func GetUserId(ctx *gin.Context) string {
	payload, err := Get(ctx)
	if err != nil {
		panic(result.NoAuth)
	}

	return payload.UserId
}

// Get 获取当前用户令牌
func Get(ctx *gin.Context) (payload *Payload, err error) {
	var token string

	// 从 header 中获取令牌
	getHeaderToken := func() string {
		if token = ctx.GetHeader("Authorization"); token == "" {
			return ""
		}

		prefix := "Bearer "
		if strings.Index(token, prefix) != 0 {
			return ""
		}

		return token[len(prefix):]
	}

	// 从 cookie 中获取令牌
	if token = getHeaderToken(); token == "" {
		if token, err = ctx.Cookie(tokenKey); err != nil {
			return nil, err
		}
	}

	return decode(token)
}

// Remove 移除令牌
func Remove(ctx *gin.Context) {
	ctx.SetCookie("token", "", -1, "/", "localhost", false, true)
}
