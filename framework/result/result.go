/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/01
 */
package result

import (
	"encoding/json"
)

// 结果对象
type Result struct {
	Code    string      `json:"code"`    // 代码
	Message string      `json:"message"` // 说明
	Data    interface{} `json:"data"`    // 数据
}

// 设置结果对象参数对
func (result Result) WithKeyPair(key string, value interface{}) Result {
	dataMap := make(map[string]interface{})
	dataMap[key] = value
	result.Data = dataMap
	return result
}

// 设置结果对象数据
func (result Result) WithData(data interface{}) Result {
	result.Data = data
	return result
}

// 设置错误对象数据
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

// 创建结果对象
func createResult(code string, message string) Result {
	return Result{
		Code:    code,
		Message: message,
	}
}

// 通用返回对象码
var (
	Ok               = createResult("Ok", "成功")
	NoAuth           = createResult("NoAuth", "没有登录")
	Duplicate        = createResult("Duplicate", "已经存在")
	Forbidden        = createResult("Forbidden", "没有权限")
	NotFound         = createResult("NotFound", "不存在")
	NotMatch         = createResult("NotMatch", "不匹配")
	RateLimit        = createResult("ExceedLimit", "超出限制")
	LogicError       = createResult("LogicError", "逻辑错误")
	ParameterError   = createResult("ParameterError", "参数错误")
	MethodNotAllowed = createResult("MethodNotAllowed", "请求方式不允许")
	InternalError    = createResult("InternalError", "内部错误")
	ThirdPartError   = createResult("ThirdPartError", "第三方错误")
)
