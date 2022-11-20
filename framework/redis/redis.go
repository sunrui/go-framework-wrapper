/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-04-16 22:43:52
 */

package redis

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

// Set 设置对象
func (rediz *Redis) Set(key string, value any, expired time.Duration) {
	pool := rediz.Pool.Get()
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

	if _, err := pool.Do("SET", key, value, "EX", fmt.Sprintf("%d", expired/time.Second)); err != nil {
		panic(err.Error())
	}
}

// GetString 获取字符串
func (rediz *Redis) GetString(key string) *string {
	pool := rediz.Pool.Get()
	defer func() {
		_ = pool.Close()
	}()

	if reply, err := pool.Do("GET", key); err != nil {
		panic(err.Error())
	} else if reply == nil {
		return nil
	} else {
		dst := fmt.Sprintf("%s", reply)
		return &dst
	}
}

// GetJson 获取对象
func (rediz *Redis) GetJson(key string, dst any) bool {
	pool := rediz.Pool.Get()
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
func (rediz *Redis) Exists(key string) bool {
	pool := rediz.Pool.Get()
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
func (rediz *Redis) Del(key string) {
	pool := rediz.Pool.Get()
	defer func() {
		_ = pool.Close()
	}()

	if _, err := pool.Do("DEL", key); err != nil {
		panic(err.Error())
	}
}

// SetHash 设置 hash 对象
func (rediz *Redis) SetHash(hash string, key string, value any) {
	pool := rediz.Pool.Get()
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
func (rediz *Redis) GetHash(hash string, key string) *string {
	pool := rediz.Pool.Get()
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
func (rediz *Redis) GetHashJson(hash string, key string, dst any) bool {
	pool := rediz.Pool.Get()
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
