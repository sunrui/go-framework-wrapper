/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-04-25 20:36:20
 */

package result

import (
	"encoding/json"
)

// Result 结果对象
type Result struct {
	Code    int         `json:"code"`           // 状态码
	Status  string      `json:"status"`         // 结果
	Message string      `json:"message"`        // 消息
	Data    interface{} `json:"data,omitempty"` // 数据
}

// Pagination 分页对象
type Pagination struct {
	Page      int   `json:"page"`      // 当前页，从 0 开始
	PageSize  int   `json:"pageSize"`  // 每页大小
	TotalPage int64 `json:"totalPage"` // 总页数
	TotalSize int64 `json:"totalSize"` // 总大小
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

// WithIdData 设置 id 结果对象数据
func (result Result) WithIdData(id string) Result {
	type idData struct {
		Id string `json:"id"`
	}

	result.Data = idData{
		Id: id,
	}
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
	Ok               = newResult(200, "Ok", "成功")
	BadRequest       = newResult(400, "BadRequest", "语法错误")
	NoAuth           = newResult(401, "NoAuth", "没有登录")
	Forbidden        = newResult(403, "Forbidden", "没有权限")
	NotFound         = newResult(404, "NotFound", "不存在")
	MethodNotAllowed = newResult(405, "MethodNotAllowed", "请求方式不允许")
	Conflict         = newResult(409, "Conflict", "请求冲突")
	RateLimit        = newResult(429, "RateLimit", "限流")
	InternalError    = newResult(500, "InternalError", "内部错误")
	NotImplemented   = newResult(501, "NotImplemented", "未实现")
	BadGateway       = newResult(502, "BadGateway", "网关错误")
)

// 创建结果对象
func newResult(code int, status string, message string) Result {
	return Result{
		Status:  status,
		Code:    code,
		Message: message,
	}
}

// All 获取所有 result 对象
func All() []Result {
	return []Result{
		Ok,
		BadRequest, NoAuth, Forbidden, NotFound, MethodNotAllowed, Conflict, RateLimit,
		InternalError, NotImplemented, BadGateway,
	}
}
