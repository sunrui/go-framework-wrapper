/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-04-25 22:14:14
 */

package area

import (
	"framework/proto/result"
	"github.com/gin-gonic/gin"
	"service/area"
	"strconv"
)

// 获取市
func getCity(ctx *gin.Context) result.Result {
	provinceId, err := strconv.Atoi(ctx.Param("provinceId"))
	if err != nil {
		return result.ParameterBindError.WithKeyPair("provinceId", ctx.Param("provinceId"))
	}

	cities := area.GetCity(provinceId)
	if cities == nil {
		return result.NotFound.WithKeyPair("provinceId", provinceId)
	}

	return result.Ok.WithData(cities)
}
