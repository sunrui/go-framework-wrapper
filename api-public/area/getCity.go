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

// 获取市
func getCity(ctx *gin.Context) {
	provinceId, err := strconv.Atoi(ctx.Param("provinceId"))
	if err != nil {
		response.New(ctx).Result(result.BadRequest.WithKeyPair("provinceId", ctx.Param("provinceId")))
		return
	}

	cities := area.GetCity(provinceId)
	if cities == nil {
		response.New(ctx).Result(result.NotFound.WithKeyPair("provinceId", provinceId))
		return
	}

	response.New(ctx).Data(cities)
}
