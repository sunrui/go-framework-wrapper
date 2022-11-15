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

// CodeType 代码类型
type CodeType string

const (
	OK               CodeType = "OK"
	NoAuth           CodeType = "NoAuth"
	ParameterError   CodeType = "ParameterError"
	Forbidden        CodeType = "Forbidden"
	NotFound         CodeType = "NotFound"
	NoMatch          CodeType = "NoMatch"
	NoContent        CodeType = "NoContent"
	MethodNotAllowed CodeType = "MethodNotAllowed"
	Conflict         CodeType = "Conflict"
	RateLimit        CodeType = "RateLimit"
	InternalError    CodeType = "InternalError"
	ThirdPartyError  CodeType = "ThirdPartyError"
	NotImplemented   CodeType = "NotImplemented"
)

// Result 结果对象
type Result[T any] struct {
	Code       CodeType         `json:"code" example:"Ok"`              // 结果
	Message    string           `json:"message,omitempty" example:"成功"` // 消息
	Data       T                `json:"data,omitempty"`                 // 数据
	Pagination *Pagination      `json:"pagination,omitempty"`           // 分页对象
	Request    *request.Request `json:"request,omitempty"`              // 请求对象
}

// String 数据
func (result Result[T]) String() string {
	marshal, _ := json.MarshalIndent(result, "", "\t")
	return string(marshal)
}

// KeyValueData 键值数据
func KeyValueData(key string, value any) map[string]any {
	dataMap := make(map[string]any)
	dataMap[key] = value
	return dataMap
}

// IdData Id 数据
func IdData(id string) any {
	type idData struct {
		Id string `json:"id"`
	}

	return idData{
		Id: id,
	}
}
