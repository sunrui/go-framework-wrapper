/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-10-15 09:11:59
 */

package result

import (
	"encoding/json"
	"framework/app/request"
)

type M map[string]any

// Pagination 分页对象
type Pagination struct {
	Page      int   `json:"page"`      // 当前页，从 1 开始
	PageSize  int   `json:"pageSize"`  // 每页大小
	TotalPage int64 `json:"totalPage"` // 总页数
	TotalSize int64 `json:"totalSize"` // 总大小
}

// Result 结果对象
type Result struct {
	Code       string           `json:"code" example:"Ok"`    // 代码
	Message    string           `json:"message" example:"成功"` // 消息
	Data       any              `json:"data,omitempty"`       // 数据
	Pagination *Pagination      `json:"pagination,omitempty"` // 分页对象
	Request    *request.Request `json:"request,omitempty"`    // 请求对象
}

// String 数据
func (result Result) String() string {
	marshal, _ := json.MarshalIndent(result, "", "\t")
	return string(marshal)
}

func (result *Result) WithMessage(message string) *Result {
	result.Message = message
	return result
}

func (result *Result) WithRequest(request request.Request) *Result {
	result.Request = &request
	return result
}

// WithData 设置 data
func (result *Result) WithData(data any) *Result {
	result.Data = data
	return result
}

// WithData 设置 data 和 pagination
func (result *Result) WithDataAndPagination(data any, pagination *Pagination) *Result {
	result.Data = data
	result.Pagination = pagination
	return result
}

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
