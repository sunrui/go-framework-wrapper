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
)

// Token 令牌对象
type Token struct {
	UserId string `json:"userId"`
}

// claims 对象
type claims struct {
	jwt.StandardClaims
	Token
}

// 密钥
var secret = config.Jwt().Secret

// 过期时间
var maxAge = config.Jwt().MaxAge

// 令牌名称
const name = "token"

// 生成 jwt 字符串
func encode(token Token) (string, error) {
	claims := claims{
		jwt.StandardClaims{},
		token,
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenClaims.SignedString(secret)
}

// 验证 jwt 字符串
func decode(tokenString string) (*Token, error) {
	tokenClaims, err := jwt.ParseWithClaims(tokenString, &claims{}, func(token *jwt.Token) (any, error) {
		return secret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*claims); ok && tokenClaims.Valid {
			return &claims.Token, nil
		}
	}

	return nil, err
}

// Write 写入 cookie 令牌
func Write(ctx *gin.Context, userId string) {
	// 生成用户令牌
	if token, err := encode(Token{
		UserId: userId,
	}); err != nil {
		panic(err.Error())
	} else {
		// 写入令牌，默认 30 天
		ctx.SetCookie(name, token, maxAge, "/", "", false, true)
	}
}

// MustGetUserId 获取当前用户 id
func MustGetUserId(ctx *gin.Context) string {
	if token, err := GetToken(ctx); err != nil {
		panic(result.NoAuth)
	} else {
		return token.UserId
	}
}

// GetToken 获取当前用户令牌
func GetToken(ctx *gin.Context) (*Token, error) {
	var tokenString string
	var err error

	// 从 cookie 中获取令牌
	if tokenString, err = ctx.Cookie(name); err == nil {
		return decode(tokenString)
	}

	// 从 header 中获取令牌
	if tokenString = ctx.GetHeader(name); tokenString != "" {
		return decode(tokenString)
	}

	// 从 Authorization 中获取令牌
	if tokenString = ctx.GetHeader("Authorization"); tokenString != "" {
		prefix := "Bearer "
		if strings.Index(tokenString, prefix) == 0 {
			tokenString = tokenString[len(prefix):]
			return decode(tokenString)
		}
	}

	// 从 cookie 中获取令牌
	return nil, errors.New("<null>")
}

// Remove 移除令牌
func Remove(ctx *gin.Context) result.Result {
	ctx.SetCookie(name, "", -1, "/", "localhost", false, true)
	return result.Ok
}
