/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/20 21:16:20
 */

package sms

import (
	"fmt"
	"medium-server-go/enum"
	"medium-server-go/framework/db"
)

// CodeCache 缓存数据
type CodeCache struct {
	Code      string `json:"code"`      // 验证码
	ErrVerify int    `json:"errVerify"` // 出错较验次数
}

// Cache 缓存对象
type Cache struct {
	Phone    string        `json:"phone"`    // 手机号
	CodeType enum.CodeType `json:"codeType"` // 验证码类型
}

// 获取主键
func (cache *Cache) getKey() string {
	return fmt.Sprintf("SMS_%s_%s", cache.CodeType, cache.Phone)
}

// 获取缓存的值
func (cache *Cache) getValue() *CodeCache {
	var codeCache CodeCache

	err := db.Redis.GetJson(cache.getKey(), &codeCache)
	if err == nil {
		return &codeCache
	}

	return nil
}

// Exists 获取缓存是否存在
func (cache *Cache) Exists() bool {
	return db.Redis.Exists(cache.getKey())
}

// Save 设置新缓存验证码
func (cache *Cache) Save(codeCache CodeCache) {
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
		value.ErrVerify += 1

		// 如果已经较验出错 5 次，移除现有验证码
		if value.ErrVerify == 5 {
			cache.Del()
		} else {
			// 更新出错较验次数
			cache.Save(*value)
		}

		return false
	}

	return true
}
