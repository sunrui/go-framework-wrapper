/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-04-16 22:43:52
 */

package db

import (
	"encoding/json"
	"fmt"
	"framework/config"
	"github.com/garyburd/redigo/redis"
	"reflect"
	"time"
)

// Redis 数据库访问对象
type redisPool struct {
	redis.Pool
}

// Redis 对象
var Redis *redisPool

// 初始化
func init() {
	// 建立连接池
	Redis = &redisPool{
		redis.Pool{
			MaxIdle:     5,
			MaxActive:   100,
			IdleTimeout: 1 * time.Hour,
			Wait:        true,
			Dial: func() (redis.Conn, error) {
				address := fmt.Sprintf("%s:%d", config.Get().Redis.Host, config.Get().Redis.Port)
				timeout := time.Duration(config.Get().Redis.Timeout) * time.Second

				return redis.Dial("tcp", address,
					redis.DialPassword(config.Get().Redis.Password),
					redis.DialDatabase(config.Get().Redis.Database),
					redis.DialConnectTimeout(timeout),
					redis.DialReadTimeout(timeout),
					redis.DialWriteTimeout(timeout))
			},
		},
	}

	// 尝试数据库连接
	if _, err := Redis.Get().Do("PING"); err != nil {
		panic(err.Error())
	}
}

// Set 设置对象
func (redisPool *redisPool) Set(key string, value any, second time.Duration) {
	pool := redisPool.Get()
	defer func() {
		_ = pool.Close()
	}()

	// 判断存储的是否为对象
	if reflect.TypeOf(value).Kind() == reflect.Struct {
		if marshal, err := json.Marshal(value); err != nil {
			panic(err.Error())
		} else {
			value = string(marshal)
		}
	}

	if _, err := pool.Do("SET", key, value, "EX", fmt.Sprintf("%d", second)); err != nil {
		panic(err.Error())
	}
}

// GetString 获取字符串
func (redisPool *redisPool) GetString(key string) *string {
	pool := redisPool.Get()
	defer func() {
		_ = pool.Close()
	}()

	if reply, err := pool.Do("GET", key); err != nil {
		panic(err.Error())
	} else if reply == nil {
		return nil
	} else {
		replyString := fmt.Sprintf("%s", reply)
		return &replyString
	}
}

// GetJson 获取对象
func (redisPool *redisPool) GetJson(key string, dst any) bool {
	pool := redisPool.Get()
	defer func() {
		_ = pool.Close()
	}()

	if reply, err := pool.Do("GET", key); err != nil {
		panic(err.Error())
	} else if reply == nil {
		return false
	} else if err = json.Unmarshal(reply.([]uint8), dst); err != nil {
		panic(err.Error())
	} else {
		return true
	}
}

// Exists 是否存在对象
func (redisPool *redisPool) Exists(key string) bool {
	pool := redisPool.Get()
	defer func() {
		_ = pool.Close()
	}()

	if ret, err := pool.Do("EXISTS", key); err != nil {
		panic(err.Error())
	} else {
		return ret.(int64) == 1
	}
}

// Del 删除对象
func (redisPool *redisPool) Del(key string) {
	pool := redisPool.Get()
	defer func() {
		_ = pool.Close()
	}()

	if _, err := pool.Do("DEL", key); err != nil {
		panic(err.Error())
	}
}

// SetHash 设置 hash 对象
func (redisPool *redisPool) SetHash(hash string, key string, value any) {
	pool := redisPool.Get()
	defer func() {
		_ = pool.Close()
	}()

	// 判断存储的是否为对象
	if reflect.TypeOf(value).Kind() == reflect.Struct {
		if marshal, err := json.Marshal(value); err != nil {
			panic(err.Error())
		} else {
			value = string(marshal)
		}
	}

	if _, err := pool.Do("HSET", hash, key, value); err != nil {
		panic(err.Error())
	}
}

// GetHash 获取 hash 对象
func (redisPool *redisPool) GetHash(hash string, key string) *string {
	pool := redisPool.Get()
	defer func() {
		_ = pool.Close()
	}()

	if reply, err := pool.Do("HGET", hash, key); err != nil {
		panic(err.Error())
	} else if reply == nil {
		return nil
	} else {
		replyString := fmt.Sprintf("%s", reply)
		return &replyString
	}
}

// GetHashJson 获取 hash 对象
func (redisPool *redisPool) GetHashJson(hash string, key string, dst any) bool {
	pool := redisPool.Get()
	defer func() {
		_ = pool.Close()
	}()

	if reply, err := pool.Do("HGET", hash, key); err != nil {
		panic(err.Error())
	} else if reply == nil {
		return false
	} else if err = json.Unmarshal(reply.([]uint8), dst); err != nil {
		panic(err.Error())
	} else {
		return true
	}
}
