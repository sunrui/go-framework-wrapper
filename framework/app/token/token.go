/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-10-15 11:12:33
 */

package token

import (
	"config"
	"errors"
	"framework/result"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"strings"
	"time"
)

// Payload 自定义对象
type Payload struct {
	UserId string `json:"userId"` // 用户 id
}

// 令牌对象
type token struct {
	jwt.StandardClaims
	Payload
}

// 生成 jwt 字符串
func encode(payload Payload) (string, error) {
	token := token{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + config.Inst().Token.MaxAge*1000,
		},
		payload,
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, token)
	return tokenClaims.SignedString(config.Inst().Token.JwtSecret)
}

// 验证 jwt 字符串
func decode(tokenString string) (*token, error) {
	tokenClaims, err := jwt.ParseWithClaims(tokenString, &token{}, func(token *jwt.Token) (any, error) {
		return config.Inst().Token.JwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*token); ok && tokenClaims.Valid {
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
		ctx.SetCookie(config.Inst().Token.Key, token, int(config.Inst().Token.MaxAge), "/", "", false, true)
	}
}

// 获取当前用户令牌
func get(ctx *gin.Context) (*token, error) {
	var tokenString string
	var err error

	// 从 cookie 中获取令牌
	if tokenString, err = ctx.Cookie(config.Inst().Token.Key); err != nil {
		// 从 header 中获取令牌
		if tokenString = ctx.GetHeader(config.Inst().Token.Key); tokenString == "" {
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

// GetUserId 获取当前用户 id
func GetUserId(ctx *gin.Context) *string {
	if token, err := get(ctx); err != nil {
		return nil
	} else {
		return &token.UserId
	}
}

// MustGetUserId 强制获取当前用户 id
func MustGetUserId(ctx *gin.Context) string {
	if token, err := get(ctx); err != nil {
		panic(result.NoAuth)
	} else {
		return token.UserId
	}
}

// RefreshIf 刷新令牌
func RefreshIf(ctx *gin.Context) {
	if token, err := get(ctx); err == nil {
		// 当前距离过期时间（毫秒）
		InstExpired := token.ExpiresAt - time.Now().Unix()

		// 设置距离过期时间（毫秒）
		setExpired := (config.Inst().Token.MaxAge - config.Inst().Token.AutoRefreshAge) * 1000

		// 已经大于最小刷新时长
		if InstExpired >= setExpired {
			Write(ctx, token.Payload)
		}
	}
}

// Remove 移除令牌
func Remove(ctx *gin.Context) {
	ctx.SetCookie(config.Inst().Token.Key, "", -1, "/", "", false, true)
}
