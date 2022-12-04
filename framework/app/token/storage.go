/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-28 16:33:50
 */

package token

import (
	"encoding/json"
	"errors"
	"framework/app/redis"
	"framework/app/util"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"time"
)

// Storage 存储
type Storage interface {
	// Set 设置
	Set(payload Payload, maxAge int64) (value string, err error)
	// Get 获取
	Get(value string) (payload *Payload, ttl int64, err error)
}

// jwt 负荷
type jwtPayload struct {
	jwt.StandardClaims
	Payload
}

// JwtStorage 存储
type JwtStorage struct {
	ctx    *gin.Context // gin 上下文
	Secret []byte       // jwt 私钥
}

// NewJwtStorage 创建 Jwt 存储
func NewJwtStorage(secret []byte) JwtStorage {
	return JwtStorage{
		Secret: secret,
	}
}

// Set 设置
func (jwtStorage JwtStorage) Set(payload Payload, maxAge int64) (value string, err error) {
	t := jwtPayload{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + maxAge*1000,
		},
		payload,
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, t).SignedString(jwtStorage.Secret)
}

// Get 设置
func (jwtStorage JwtStorage) Get(value string) (payload *Payload, ttl int64, err error) {
	tokenClaims, err := jwt.ParseWithClaims(value, &jwtPayload{}, func(token *jwt.Token) (any, error) {
		return jwtStorage.Secret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*jwtPayload); ok && tokenClaims.Valid {
			payload = &claims.Payload
			ttl = (claims.ExpiresAt - time.Now().Unix()) / 1000
			err = nil
			return
		}
	}

	return nil, 0, err
}

// RedisStorage Redis 存储
type RedisStorage struct {
	Redis     *redis.Redis // Redis 对象
	KeyPrefix string       // 键值前缀
}

// NewRedisStorage 创建 Redis 存储
func NewRedisStorage(redis *redis.Redis, keyPrefix string) RedisStorage {
	return RedisStorage{
		Redis:     redis,
		KeyPrefix: keyPrefix,
	}
}

// Set 设置负荷
func (redisStorage RedisStorage) Set(payload Payload, maxAge int64) (value string, err error) {
	key := redisStorage.KeyPrefix + "_" + util.CreateNanoid(16)
	marshal, _ := json.MarshalIndent(payload, "", "\t")
	redisStorage.Redis.Set(key, marshal, time.Duration(maxAge))
	return key, nil
}

// Get 获取负荷
func (redisStorage RedisStorage) Get(value string) (payload *Payload, ttl int64, err error) {
	var valueString []byte
	var ok bool

	if valueString, ttl, ok = redisStorage.Redis.Get(value); !ok {
		return nil, 0, errors.New("get value failed")
	} else {
		if err = json.Unmarshal(valueString, &payload); err != nil {
			return nil, 0, err
		} else {
			return payload, ttl, nil
		}
	}
}
