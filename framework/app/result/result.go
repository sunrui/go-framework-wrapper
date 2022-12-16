/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-10-15 09:11:59
 */

package result

import (
	"encoding/json"
	"net/http"
)

// M 键值
type M map[string]any

// Pagination 分页
type Pagination struct {
	Page      int   `json:"page"`      // 当前页，从 1 开始
	PageSize  int   `json:"pageSize"`  // 每页大小
	TotalPage int64 `json:"totalPage"` // 总页数
	TotalSize int64 `json:"totalSize"` // 总大小
}

// Request 请求
type Request struct {
	Ip     string      `json:"ip"`     // ip 地址
	Method string      `json:"method"` // 请求方式
	Uri    string      `json:"uri"`    // 访问地址
	Header http.Header `json:"header"` // server 首部
	Body   *string     `json:"body"`   // 请求体
}

// Result 结果
type Result struct {
	Code       string      `json:"code" example:"Ok"`      // 代码
	Message    string      `json:"message" example:"成功"` // 消息
	Data       any         `json:"data,omitempty"`         // 数据
	Elapsed    int64       `json:"elapsed"`                // 耗时
	Pagination *Pagination `json:"pagination,omitempty"`   // 分页config
	Request    *Request    `json:"request,omitempty"`      // 请求config
}

// String 数据
func (result *Result) String() string {
	marshal, _ := json.MarshalIndent(result, "", "\t")
	return string(marshal)
}

// WithMessage 设置消息
func (result *Result) WithMessage(message string) *Result {
	result.Message = message
	return result
}

// WithRequest 设置请求
func (result *Result) WithRequest(request Request) *Result {
	result.Request = &request
	return result
}

// WithData 设置 data
func (result *Result) WithData(data any) *Result {
	result.Data = data
	return result
}

// WithDataAndPagination WithData 设置 data 和 pagination
func (result *Result) WithDataAndPagination(data any, pagination *Pagination) *Result {
	result.Data = data
	result.Pagination = pagination
	return result
}

// 创建
func newResult(code string, message string) Result {
	return Result{
		Code:    code,
		Message: message,
	}
}

var (
	Ok               = newResult("Ok", "成功")
	NoAuth           = newResult("NoAuth", "未登录")
	ParameterError   = newResult("ParameterError", "参数错误")
	Forbidden        = newResult("Forbidden", "没有权限")
	NotFound         = newResult("NotFound", "没有找到")
	NoMatch          = newResult("NoMatch", "不匹配")
	NoContent        = newResult("NoContent", "没有内容")
	MethodNotAllowed = newResult("MethodNotAllowed", "方法不允许")
	Conflict         = newResult("Conflict", "冲突")
	RateLimit        = newResult("RateLimit", "限流")
	InternalError    = newResult("InternalError", "内部错误")
	ThirdPartyError  = newResult("ThirdPartyError", "第三方错误")
	NotImplemented   = newResult("NotImplemented", "未实现")
)
