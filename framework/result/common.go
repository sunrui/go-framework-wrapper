/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2021/12/31
 */

package result

// 通用返回对象码
var (
	Ok               = createResult("Ok", "成功")
	NoAuth           = createResult("NoAuth", "没有登录")
	Duplicate        = createResult("Duplicate", "已经存在")
	Forbidden        = createResult("Forbidden", "没有权限")
	NotFound         = createResult("NotFound", "不存在")
	NotMatch         = createResult("NotMatch", "不匹配")
	RateLimit        = createResult("ExceedLimit", "超出限制")
	LogicError       = createResult("LogicError", "逻辑错误")
	ParameterError   = createResult("ParameterError", "参数错误")
	MethodNotAllowed = createResult("MethodNotAllowed", "请求方式不允许")
	InternalError    = createResult("InternalError", "内部错误")
	ThirdPartError   = createResult("ThirdPartError", "第三方错误")
)

// 创建结果对象
func createResult(code string, message string) Result {
	return Result{
		Code:    code,
		Message: message,
	}
}
