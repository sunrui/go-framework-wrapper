/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-10-15 10:22:41
 */

package app

import (
	"framework/request"
	"framework/result"
	"github.com/gin-gonic/gin"
	"net/http"
)

// validator 较验
// https://github.com/go-playground/validator/

// PageRequest 分页请求对象
type PageRequest struct {
	Page     int `json:"page" form:"page" validate:"required,gte=1,lte=9999"`       // 分页，从 1 开始
	PageSize int `json:"pageSize" form:"pageSize" validate:"required,gte=1,lte=99"` // 分页大小
}

// Reply 回复
func Reply(ctx *gin.Context, result result.Result) {
	// 是否结果导出请求
	if request.IsDebugRequest(ctx) {
		req := request.GetRequest(ctx)
		result.Request = &req
	}

	// 返回客户端
	ctx.JSON(http.StatusOK, result)

	// 记录日志
	WriteLog(ctx, result)
}
