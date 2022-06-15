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
	Code        string      `json:"code" enums:"Ok,BadRequest,NoAuth,Forbidden,NotFound,MethodNotAllowed,Conflict,RateLimit,InternalError,ThirdPartyError,NotImplemented" example:"Ok"` // 结果
	Description string      `json:"description" example:"成功"`                                                                                                                           // 描述
	Data        interface{} `json:"data,omitempty"`                                                                                                                                     // 数据
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
	result.Description = message
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

// 重写返回结果对象，使用 json 反序列化
func (result Result) String() string {
	marshal, _ := json.Marshal(result)
	return string(marshal)
}

// 通用返回对象码
var (
	Ok               = newResult("Ok", "成功")
	BadRequest       = newResult("BadRequest", "语法错误")
	NoAuth           = newResult("NoAuth", "没有登录")
	Forbidden        = newResult("Forbidden", "没有权限")
	NotFound         = newResult("NotFound", "不存在")
	MethodNotAllowed = newResult("MethodNotAllowed", "请求方式不允许")
	Conflict         = newResult("Conflict", "请求冲突")
	RateLimit        = newResult("RateLimit", "限流")
	InternalError    = newResult("InternalError", "内部错误")
	ThirdPartyError  = newResult("ThirdPartyError", "第三方错误")
	NotImplemented   = newResult("NotImplemented", "未实现")
)

// 创建结果对象
func newResult(code string, description string) Result {
	return Result{
		Code:        code,
		Description: description,
	}
}

// All 获取所有 result 对象
func All() []Result {
	return []Result{
		Ok,
		BadRequest, NoAuth, Forbidden, NotFound, MethodNotAllowed, Conflict, RateLimit,
		InternalError, ThirdPartyError, NotImplemented,
	}
}
