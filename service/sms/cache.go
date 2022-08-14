/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-04-25 21:28:29
 */

package sms

import (
	"fmt"
	"framework/config"
	"framework/db"
	"time"
)

// 缓存数据
type cachePayload struct {
	Code        string `json:"code"`        // 验证码
	VerifyTimes int    `json:"verifyTimes"` // 较验次数
}

// Cache 缓存对象
type Cache struct {
	Phone   string `json:"phone"`   // 手机号
	SmsType Type   `json:"smsType"` // 验证码类型
}

// 获取主键
func (cache *Cache) getKey() string {
	return fmt.Sprintf("SMS_%s_%s", cache.SmsType, cache.Phone)
}

// Exists 获取缓存是否存在
func (cache *Cache) Exists() bool {
	return db.Redis.Exists(cache.getKey())
}

// SaveCode 设置新缓存验证码
func (cache *Cache) SaveCode(code string) {
	db.Redis.Set(cache.getKey(), cachePayload{
		Code:        code,
		VerifyTimes: 0,
	}, time.Duration(config.Sms().MaxAge))
}

// Del 移除缓存验证码
func (cache *Cache) Del() {
	db.Redis.Del(cache.getKey())
}

// Verify 较验验证码
func (cache *Cache) Verify(code string) bool {
	// 获取缓存数据
	var payload cachePayload

	if ok := db.Redis.GetJson(cache.getKey(), &payload); !ok {
		return false
	}

	// 如果验证码较验错误
	if payload.Code != code {
		// 增加缓存引用记数
		payload.VerifyTimes += 1

		// 如果已经较验出错 maxVerifyTime 次，移除现有验证码
		if payload.VerifyTimes == config.Sms().MaxVerifyTime {
			cache.Del()
		} else {
			// 更新出错较验次数
			db.Redis.Set(cache.getKey(), payload, time.Duration(config.Sms().MaxAge))
		}

		return false
	}

	return true
}