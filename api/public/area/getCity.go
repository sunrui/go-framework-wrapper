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
	if provinceId, err := strconv.Atoi(ctx.Param("provinceId")); err != nil {
		return result.ParameterBindError.WithKeyPair("provinceId", ctx.Param("provinceId"))
	} else {
		if cities := area.GetCity(provinceId); cities == nil {
			return result.NotFound.WithKeyPair("provinceId", provinceId)
		} else {
			return result.Ok.WithData(cities)
		}
	}
}
