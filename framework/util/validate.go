/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-10-15 09:03:33
 */

package util

import (
	"errors"
	"fmt"
	"framework/result"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strings"
)

// ValidateParameter 请求参数过滤
func ValidateParameter(ctx *gin.Context, req any) {
	var validationErrors validator.ValidationErrors
	var err error

	// 绑定类型
	var bindingType = func() binding.Binding {
		if ctx.Request.Method == http.MethodGet {
			return binding.Query
		} else {
			return binding.JSON
		}
	}()

	// 强制解析
	if err = ctx.ShouldBindWith(req, bindingType); err != nil {
		panic(result.ParameterBindError.WithData(fmt.Sprintf("%s", err)))
	}

	// 存在解析参数错误
	if err = validator.New().Struct(req); err != nil {
		// 参数错误对象
		type ParamError struct {
			Field    string `json:"field"`    // 变量名
			Validate string `json:"validate"` // 较验值
		}

		// 解析内容出错
		if errors.As(err, &validationErrors) {
			var parameter []ParamError

			// 遍历解析参数
			for _, e := range validationErrors {
				validate := e.Tag()
				if len(e.Param()) != 0 {
					validate += "=" + e.Param()
				}

				parameter = append(parameter, ParamError{
					Field:    strings.ToLower(e.Field()),
					Validate: validate,
				})
			}

			panic(result.ParameterValidateError.WithData(parameter))
		}
	}
}
