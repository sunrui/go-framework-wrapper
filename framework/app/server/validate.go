/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-10-15 09:03:33
 */

package server

import (
	"errors"
	"framework/app/result"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"net/http"
)

// ValidateParameter 请求参数过滤
func ValidateParameter(ctx *gin.Context, req any) {
	var validationErrors validator.ValidationErrors
	var err error

	// 绑定类型
	var bindingType binding.Binding
	if ctx.Request.Method == http.MethodGet {
		bindingType = binding.Query
	} else {
		bindingType = binding.JSON
	}

	// 强制解析
	if err = ctx.ShouldBindWith(req, bindingType); err != nil {
		panic(result.ParameterError.WithData(err.Error()))
	}

	// 存在解析参数错误
	if err = validator.New().Struct(req); err != nil {
		// 参数错误
		type ParamError struct {
			Field string `json:"field"` // 变量名
			Error string `json:"error"` // 较验值
		}

		// 解析内容出错
		if errors.As(err, &validationErrors) {
			var parameter []ParamError

			// 遍历解析参数
			for _, e := range validationErrors {
				parameter = append(parameter, ParamError{
					Field: e.Field(),
					Error: e.Error(),
				})
			}

			panic(result.ParameterError.WithData(parameter))
		}
	}
}
