/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-06-15 15:22:23
 */

package _sample

import (
	"errors"
	"fmt"
	"framework/app"
	"framework/proto/response"
	"framework/proto/result"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strings"
)

func getSample(ctx *gin.Context) {
	// 分页请求对象
	var req getSampleReq

	var validationErrors validator.ValidationErrors
	var bindingType binding.Binding
	var err error

	if ctx.Request.Method == http.MethodGet {
		bindingType = binding.Query
	} else {
		bindingType = binding.JSON
	}

	// 强制解析
	if err = ctx.ShouldBindWith(req, bindingType); err != nil {
		goto ERROR
	}

	// 存在解析参数错误
	if err = validator.New().Struct(req); err != nil {
		goto ERROR
	}

	return

ERROR:
	// 参数错误对象
	type ParamError struct {
		Field    string `json:"field"`    // 变量名
		Validate string `json:"validate"` // 较验值
	}

	dataMap := make(map[string]interface{})

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

		dataMap["parameter"] = parameter
	} else {
		dataMap["error"] = fmt.Sprintf("%s", err)
	}

	panic(result.BadRequest.WithData(dataMap))

	// 较验参数
	app.ValidateParameter(ctx, &req)

	println(req.PageSize)

	response.New(ctx).Data("hello world")
}

func putSample(ctx *gin.Context) {
	response.New(ctx).Data("hello world")
}
