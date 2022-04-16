/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022/04/16 15:19:16
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
