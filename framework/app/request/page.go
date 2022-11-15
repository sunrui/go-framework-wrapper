/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-14 23:36:31
 */

package request

// validator 较验
// https://github.com/go-playground/validator/

// Page 分页请求对象
type Page struct {
	Page     int `json:"page" form:"page" validate:"required,gte=1,lte=9999"`       // 分页，从 1 开始
	PageSize int `json:"pageSize" form:"pageSize" validate:"required,gte=1,lte=99"` // 分页大小
}
