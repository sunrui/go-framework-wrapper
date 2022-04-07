/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-04-07 16:12:11
 */

package template

// 创建请求
type postTemplateReq struct {
	Name string `json:"name" validate:"required"` // 名称
}

// 更新请求
type putTemplateReq struct {
	Name string `json:"name" validate:"required"` // 名称
}
