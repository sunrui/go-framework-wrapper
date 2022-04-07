/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/29 18:02:29
 */

package area

import (
	"framework/proto/response"
	"framework/proto/result"
	"github.com/gin-gonic/gin"
	"service/core/area"
	"strconv"
)

// 获取国家
func getCountry(ctx *gin.Context) {
	country := area.GetCountry()
	response.New(ctx).Data(country)
}

// 获取省
func getProvince(ctx *gin.Context) {
	provinces := area.GetProvinces()
	response.New(ctx).Data(provinces)
}

// 获取市
func getCity(ctx *gin.Context) {
	provinceId, err := strconv.Atoi(ctx.Param("provinceId"))
	if err != nil {
		response.New(ctx).Data(result.ParameterError.WithKeyPair("provinceId", ctx.Param("provinceId")))
		return
	}

	cities := area.GetCity(provinceId)
	if cities == nil {
		response.New(ctx).Data(result.NotFound.WithKeyPair("provinceId", provinceId))
		return
	}

	response.New(ctx).Data(cities)
}

// 获取地区
func getArea(ctx *gin.Context) {
	cityId, err := strconv.Atoi(ctx.Param("cityId"))
	if err != nil {
		response.New(ctx).Data(result.ParameterError.WithKeyPair("cityId", ctx.Param("cityId")))
		return
	}

	areas := area.GetArea(cityId)
	if areas == nil {
		response.New(ctx).Data(result.NotFound.WithKeyPair("cityId", cityId))
		return
	}

	response.New(ctx).Data(areas)
}
