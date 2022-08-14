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

// 获取地区
func getArea(ctx *gin.Context) result.Result {
	cityId, err := strconv.Atoi(ctx.Param("cityId"))
	if err != nil {
		return result.ParameterBindError.WithKeyPair("cityId", ctx.Param("cityId"))
	}

	areas := area.GetArea(cityId)
	if areas == nil {
		return result.NotFound.WithKeyPair("cityId", cityId)
	}

	return result.Ok.WithData(areas)
}
