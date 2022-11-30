/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-12-01 04:59:39
 */

package context

import "framework/app/token"

// Tokens 令牌
type Tokens struct {
	Jwt   *token.Token // jwt
	Redis *token.Token // redis
}

// 创建令牌
func newTokens() *Tokens {
	return &Tokens{
		Jwt: func() *token.Token {
			jwtStorage := token.NewJwtStorage([]byte(Config.Token.JwtSecret))
			return token.New(Config.Token, jwtStorage)
		}(),
		Redis: func() *token.Token {
			redisStorage := token.NewRedisStorage(Redis, Config.Token.Key)
			return token.New(Config.Token, redisStorage)
		}(),
	}
}
