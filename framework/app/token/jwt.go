/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-26 19:38:34
 */

package token

import (
	"github.com/golang-jwt/jwt"
	"time"
)

// Payload 负荷
type Payload struct {
	UserId string `json:"userId"` // 用户 id
}

// jwt 负荷
type jwtPayload struct {
	jwt.StandardClaims
	Payload
}

// 生成 jwt 字符串
func encode(payload Payload, maxAge int64, jwtSecret string) (string, error) {
	t := jwtPayload{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + maxAge*1000,
		},
		payload,
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, t)
	return tokenClaims.SignedString(jwtSecret)
}

// 验证 jwt 字符串
func decode(tokenString string, jwtSecret string) (*jwtPayload, error) {
	tokenClaims, err := jwt.ParseWithClaims(tokenString, &jwtPayload{}, func(token *jwt.Token) (any, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*jwtPayload); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
