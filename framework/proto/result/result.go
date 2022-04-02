/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/01
 */

package result

import (
	"encoding/json"
	"net/http"
)

// Result 结果对象
type Result struct {
	Status  int         `json:"-"`              // Http 状态码
	Code    string      `json:"code"`           // 代码
	Message string      `json:"message"`        // 消息
	Data    interface{} `json:"data,omitempty"` // 数据
}

// Pagination 分页对象
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

// WithMessage 设置消息
func (result Result) WithMessage(message string) Result {
	result.Message = message
	return result
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
	// 操作成功 200
	Ok = newResult(http.StatusOK, "Ok", "成功")

	// 业务级错误 300
	NotFound   = newResult(http.StatusMultipleChoices, "NotFound", "不存在")
	NotMatch   = newResult(http.StatusMultipleChoices, "NotMatch", "不匹配")
	Duplicate  = newResult(http.StatusMultipleChoices, "Duplicate", "重复操作")
	StateError = newResult(http.StatusMultipleChoices, "StateError", "状态错误")

	// 应用级错误 400
	NoAuth    = newResult(http.StatusBadRequest, "NoAuth", "没有登录")
	Forbidden = newResult(http.StatusBadRequest, "Forbidden", "没有权限")
	RateLimit = newResult(http.StatusBadRequest, "RateLimit", "超出限制")

	// 系统级错误 500
	MethodNotAllowed = newResult(http.StatusInternalServerError, "MethodNotAllowed", "请求方式不允许")
	ParameterError   = newResult(http.StatusInternalServerError, "ParameterError", "参数错误")
	InternalError    = newResult(http.StatusInternalServerError, "InternalError", "内部错误")
	ThirdPartyError  = newResult(http.StatusInternalServerError, "ThirdPartyError", "第三方错误")
)

// 创建结果对象
func newResult(status int, code string, message string) Result {
	return Result{
		Status:  status,
		Code:    code,
		Message: message,
	}
}
