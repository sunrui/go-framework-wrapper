/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/01
 */

package exception

import (
	"encoding/json"
)

// Exception 结果对象
type Exception struct {
	Code    string      `json:"code"`    // 代码
	Message string      `json:"message"` // 说明
	Data    interface{} `json:"data"`    // 数据
}

// WithKeyPair 设置结果对象参数对
func (exception Exception) WithKeyPair(key string, value interface{}) Exception {
	dataMap := make(map[string]interface{})
	dataMap[key] = value
	exception.Data = dataMap
	return exception
}

// WithData 设置结果对象数据
func (exception Exception) WithData(data interface{}) Exception {
	exception.Data = data
	return exception
}

// WithError 设置错误对象数据
func (exception Exception) WithError(data interface{}) Exception {
	dataMap := make(map[string]interface{})
	dataMap["error"] = data
	exception.Data = dataMap
	return exception
}

// 重写返回结果对象，使用 json 反序列化
func (exception Exception) String() string {
	marshal, _ := json.Marshal(exception)
	return string(marshal)
}

// 通用返回对象码
var (
	NoAuth           = createException("NoAuth", "没有登录")
	Duplicate        = createException("Duplicate", "已经存在")
	Forbidden        = createException("Forbidden", "没有权限")
	NotFound         = createException("NotFound", "不存在")
	NotMatch         = createException("NotMatch", "不匹配")
	RateLimit        = createException("ExceedLimit", "超出限制")
	LogicError       = createException("LogicError", "逻辑错误")
	ParameterError   = createException("ParameterError", "参数错误")
	MethodNotAllowed = createException("MethodNotAllowed", "请求方式不允许")
	InternalError    = createException("InternalError", "内部错误")
	ThirdPartyError  = createException("ThirdPartyError", "第三方错误")
)

// 创建结果对象
func createException(code string, message string) Exception {
	return Exception{
		Code:    code,
		Message: message,
	}
}
