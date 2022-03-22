/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/01
 */

package result

import (
	"encoding/json"
)

// Result 结果对象
type Result struct {
	Code    string      `json:"code"`    // 代码
	Message string      `json:"message"` // 说明
	Data    interface{} `json:"data"`    // 数据
}

// WithKeyPair 设置结果对象参数对
func (result Result) WithKeyPair(key string, value interface{}) Result {
	dataMap := make(map[string]interface{})
	dataMap[key] = value
	result.Data = dataMap
	return result
}

// WithData 设置结果对象数据
func (result Result) WithData(data interface{}) Result {
	result.Data = data
	return result
}

// WithError 设置错误对象数据
func (result Result) WithError(data interface{}) Result {
	dataMap := make(map[string]interface{})
	dataMap["error"] = data
	result.Data = dataMap
	return result
}

// 重写返回结果对象，使用 json 反序列化
func (result Result) String() string {
	marshal, _ := json.Marshal(result)
	return string(marshal)
}
