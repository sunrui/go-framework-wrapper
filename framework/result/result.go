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

// 分页对象
type Pagination struct {
	Page      int `json:"page"`      // 当前页，从 0 开始
	PageSize  int `json:"pageSize"`  // 每页大小
	TotalPage int `json:"totalPage"` // 总页数
	TotalSize int `json:"totalSize"` // 总大小
}

// PageResult 结果对象
type PageResult struct {
	Result                // 结果对象
	Pagination Pagination `json:"pagination"` // 分页对象
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

// 通用返回对象码
var (
	Ok               = newResult("Ok", "成功")
	NoAuth           = newResult("NoAuth", "没有登录")
	Duplicate        = newResult("Duplicate", "重复")
	Forbidden        = newResult("Forbidden", "没有权限")
	NotFound         = newResult("NotFound", "不存在")
	RateLimit        = newResult("ExceedLimit", "超出限制")
	StateError       = newResult("StateError", "状态错误")
	ParameterError   = newResult("ParameterError", "参数错误")
	MethodNotAllowed = newResult("MethodNotAllowed", "请求方式不允许")
	InternalError    = newResult("InternalError", "内部错误")
	ThirdPartyError  = newResult("ThirdPartyError", "第三方错误")
)

// 创建结果对象
func newResult(code string, message string) Result {
	return Result{
		Code:    code,
		Message: message,
	}
}
