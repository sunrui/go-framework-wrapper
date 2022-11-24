/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-10-15 09:11:59
 */

package result

var (
	NoAuth           = newResult[any]("NoAuth", "未登录")
	ParameterError   = newResult[any]("ParameterError", "参数错误")
	Forbidden        = newResult[any]("Forbidden", "没有权限")
	NotFound         = newResult[any]("NotFound", "没有找到")
	NoMatch          = newResult[any]("NoMatch", "不匹配")
	NoContent        = newResult[any]("NoContent", "没有内容")
	MethodNotAllowed = newResult[any]("MethodNotAllowed", "方法不允许")
	Conflict         = newResult[any]("Conflict", "冲突")
	RateLimit        = newResult[any]("RateLimit", "限流")
	InternalError    = newResult[any]("InternalError", "内部错误")
	ThirdPartyError  = newResult[any]("ThirdPartyError", "第三方错误")
	NotImplemented   = newResult[any]("NotImplemented", "未实现")
)
