/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-04-16 22:43:52
 */

package redis

import (
	"fmt"
	"framework/config"
	"github.com/gomodule/redigo/redis"
	"time"
)

// Redis 缓存
type Redis struct {
	Pool redis.Pool
}

// New 创建
func New(redisConfig config.Redis) *Redis {
	// 建立连接池
	rediz := &Redis{
		Pool: redis.Pool{
			MaxIdle:     5,
			MaxActive:   100,
			IdleTimeout: 1 * time.Hour,
			Wait:        true,
			Dial: func() (redis.Conn, error) {
				address := fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port)
				timeout := time.Duration(10) * time.Second

				return redis.Dial("tcp", address,
					redis.DialPassword(redisConfig.Password),
					redis.DialDatabase(redisConfig.Database),
					redis.DialConnectTimeout(timeout),
					redis.DialReadTimeout(timeout),
					redis.DialWriteTimeout(timeout))
			},
		},
	}

	// 尝试数据库连接
	if _, err := rediz.Pool.Get().Do("PING"); err != nil {
		panic(err.Error())
	}

	return rediz
}

func (rediz *Redis) getTtl(key string) (ttl int64, ok bool) {
	pool := rediz.Pool.Get()
	defer func() {
		_ = pool.Close()
	}()

	if reply, err := pool.Do("TTL", key); err != nil {
		panic(err.Error())
	} else if reply.(int64) <= 0 {
		return 0, false
	} else {
		return reply.(int64), true
	}
}

// Set 设置
func (rediz *Redis) Set(key string, value string, expired time.Duration) {
	pool := rediz.Pool.Get()
	defer func() {
		_ = pool.Close()
	}()

	if _, err := pool.Do("SET", key, value, "EX", fmt.Sprintf("%d", expired)); err != nil {
		panic(err.Error())
	}
}

// GetString 获取字符串
func (rediz *Redis) GetString(key string) (value *string, ttl int64, ok bool) {
	pool := rediz.Pool.Get()
	defer func() {
		_ = pool.Close()
	}()

	if reply, err := pool.Do("GET", key); err != nil {
		panic(err.Error())
	} else if reply == nil {
		return nil, 0, false
	} else {
		dst := fmt.Sprintf("%s", reply)
		value = &dst
		ttl, ok = rediz.getTtl(key)
		return
	}
}

// SetHash 设置 hash
func (rediz *Redis) SetHash(hash string, key string, value any) {
	pool := rediz.Pool.Get()
	defer func() {
		_ = pool.Close()
	}()

	if _, err := pool.Do("HSET", hash, key, value); err != nil {
		panic(err.Error())
	}
}

// GetHash 获取 hash
func (rediz *Redis) GetHash(hash string, key string) (value *string, ok bool) {
	pool := rediz.Pool.Get()
	defer func() {
		_ = pool.Close()
	}()

	if reply, err := pool.Do("HGET", hash, key); err != nil {
		panic(err.Error())
	} else if reply == nil {
		return nil, false
	} else {
		dst := fmt.Sprintf("%s", reply)
		return &dst, true
	}
}

// Exists 是否存在
func (rediz *Redis) Exists(key string) bool {
	pool := rediz.Pool.Get()
	defer func() {
		_ = pool.Close()
	}()

	if reply, err := pool.Do("EXISTS", key); err != nil {
		panic(err.Error())
	} else {
		return reply.(int64) == 1
	}
}

// Del 删除
func (rediz *Redis) Del(key string) bool {
	pool := rediz.Pool.Get()
	defer func() {
		_ = pool.Close()
	}()

	if reply, err := pool.Do("DEL", key); err != nil {
		panic(err.Error())
	} else {
		return reply.(int64) == 1
	}
}
