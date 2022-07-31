/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-04-25 22:14:14
 */

package area

import (
	"framework/proto/response"
	"framework/proto/result"
	"github.com/gin-gonic/gin"
	"service/core/area"
	"strconv"
)

// 获取地区
func getArea(ctx *gin.Context) {
	cityId, err := strconv.Atoi(ctx.Param("cityId"))
	if err != nil {
		response.New(ctx).Result(result.ParameterBindError.WithKeyPair("cityId", ctx.Param("cityId")))
		return
	}

	areas := area.GetArea(cityId)
	if areas == nil {
		response.New(ctx).Result(result.NotFound.WithKeyPair("cityId", cityId))
		return
	}

	response.New(ctx).Data(areas)
}
