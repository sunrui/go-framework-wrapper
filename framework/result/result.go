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

// Pagination 分页对象
type Pagination struct {
	Page      int   `json:"page"`      // 当前页，从 1 开始
	PageSize  int   `json:"pageSize"`  // 每页大小
	TotalPage int64 `json:"totalPage"` // 总页数
	TotalSize int64 `json:"totalSize"` // 总大小
}

// Result 结果对象
type Result[T any] struct {
	Code       string           `json:"code" example:"Ok"`    // 代码
	Message    string           `json:"message" example:"成功"` // 消息
	Data       T                `json:"data,omitempty"`       // 数据
	Pagination *Pagination      `json:"pagination,omitempty"` // 分页对象
	Request    *request.Request `json:"request,omitempty"`    // 请求对象
}

// String 数据
func (result Result[T]) String() string {
	marshal, _ := json.MarshalIndent(result, "", "\t")
	return string(marshal)
}

func (result *Result[T]) WithMessage(message string) *Result[T] {
	result.Message = message
	return result
}

func (result *Result[T]) WithRequest(request request.Request) *Result[T] {
	result.Request = &request
	return result
}

func newResult[T any](code string, message string) Result[T] {
	return Result[T]{
		Code:    code,
		Message: message,
	}
}
