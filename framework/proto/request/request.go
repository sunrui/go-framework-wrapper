/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022/04/03 00:05:03
 */

package request

// PageRequest 分页请求对象
type PageRequest struct {
	Page     int `json:"page" form:"page" validate:"required,gte=1,lte=9999"`       // 分页，从 1 开始
	PageSize int `json:"pageSize" form:"pageSize" validate:"required,gte=1,lte=99"` // 分页大小
}
