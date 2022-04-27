/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-04-25 21:28:29
 */

package sms

import (
	"fmt"
	"framework/db"
	"service/enum"
)

// 缓存数据
type codeCache struct {
	Code        string `json:"code"`        // 验证码
	VerifyTimes int    `json:"verifyTimes"` // 较验次数
}

// Cache 缓存对象
type Cache struct {
	Phone   string       `json:"phone"`   // 手机号
	SmsType enum.SmsType `json:"SmsType"` // 验证码类型
}

// 获取主键
func (cache *Cache) getKey() string {
	return fmt.Sprintf("SMS_%s_%s", cache.SmsType, cache.Phone)
}

// 获取缓存的值
func (cache *Cache) getValue() *codeCache {
	var codeCache codeCache

	if success := db.Redis.GetJson(cache.getKey(), &codeCache); success {
		return &codeCache
	}

	return nil
}

// Exists 获取缓存是否存在
func (cache *Cache) Exists() bool {
	return db.Redis.Exists(cache.getKey())
}

// Save 设置新缓存验证码
func (cache *Cache) SaveCode(randomCode string) {
	db.Redis.Set(cache.getKey(), codeCache{
		Code:        randomCode,
		VerifyTimes: 0,
	}, 15*60)
}

// Save 设置新缓存验证码
func (cache *Cache) Save(codeCache codeCache) {
	db.Redis.Set(cache.getKey(), codeCache, 15*60)
}

// Del 移除缓存验证码
func (cache *Cache) Del() {
	db.Redis.Del(cache.getKey())
}

// Verify 较验验证码
func (cache *Cache) Verify(code string) bool {
	// 获取缓存数据
	value := cache.getValue()
	if value == nil {
		return false
	}

	// 如果验证码较验错误
	if value.Code != code {
		// 增加缓存引用记数
		value.VerifyTimes += 1

		// 如果已经较验出错 5 次，移除现有验证码
		if value.VerifyTimes == 5 {
			cache.Del()
		} else {
			// 更新出错较验次数
			cache.Save(*value)
		}

		return false
	}

	return true
}
