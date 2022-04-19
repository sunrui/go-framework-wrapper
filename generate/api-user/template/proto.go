/*
 * Copyright (c) $today.year honeysense.com All rights reserved.
 * Author: $author
 * Date: $today.format("yyyy-MM-dd HH:mm:ss")
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
