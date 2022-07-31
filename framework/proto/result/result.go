/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-04-25 20:36:20
 */

package result

import (
	"encoding/json"
)

// MessageType 消息类型
type MessageType string

const (
	MessageTypeToast  MessageType = "Toast"  // 消息提示
	MessageTypeDialog MessageType = "Dialog" // 消息弹窗
	MessageTypeIgnore MessageType = "Ignore" // 消息忽略
)

// Result 结果对象
type Result struct {
	Code        string      `json:"code" example:"Ok"`           // 结果
	Message     string      `json:"message" example:"成功"`        // 消息
	MessageType MessageType `json:"messageType" example:"Toast"` // 消息类型
	Data        any         `json:"data,omitempty"`              // 数据
}

// Pagination 分页对象
type Pagination struct {
	Page      int   `json:"page"`      // 当前页，从 1 开始
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
func (result Result) WithKeyPair(key string, value any) Result {
	dataMap := make(map[string]any)
	dataMap[key] = value
	result.Data = dataMap
	return result
}

// WithData 设置结果对象数据
func (result Result) WithData(data any) Result {
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

// 重写返回结果对象，使用 json 反序列化
func (result Result) String() string {
	marshal, _ := json.Marshal(result)
	return string(marshal)
}

// 通用返回对象码
var (
	Ok                     = newResult("Ok", "成功", MessageTypeToast)
	ParameterBindError     = newResult("ParameterBindError", "参数绑定错误", MessageTypeDialog)
	ParameterValidateError = newResult("ParameterValidateError", "参数较验错误", MessageTypeDialog)
	NoAuth                 = newResult("NoAuth", "没有登录", MessageTypeIgnore)
	Forbidden              = newResult("Forbidden", "没有权限", MessageTypeDialog)
	NotFound               = newResult("NotFound", "不存在", MessageTypeDialog)
	NoContent              = newResult("NoContent", "没有数据", MessageTypeDialog)
	MethodNotAllowed       = newResult("MethodNotAllowed", "请求方式不允许", MessageTypeDialog)
	Conflict               = newResult("Conflict", "请求冲突", MessageTypeDialog)
	RateLimit              = newResult("RateLimit", "限流", MessageTypeDialog)
	InternalError          = newResult("InternalError", "内部错误", MessageTypeDialog)
	ThirdPartyError        = newResult("ThirdPartyError", "第三方错误", MessageTypeDialog)
	NotImplemented         = newResult("NotImplemented", "未实现", MessageTypeDialog)
)

// 创建结果对象
func newResult(code string, message string, messageType MessageType) Result {
	return Result{
		Code:        code,
		Message:     message,
		MessageType: messageType,
	}
}

// All 获取所有 result 对象
func All() []Result {
	return []Result{
		Ok,
		ParameterBindError,
		ParameterValidateError,
		NoAuth,
		Forbidden,
		NotFound,
		NoContent,
		MethodNotAllowed,
		Conflict,
		RateLimit,
		InternalError,
		ThirdPartyError,
		NotImplemented,
	}
}
