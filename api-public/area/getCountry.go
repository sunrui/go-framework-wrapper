/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-04-25 22:14:14
 */

package area

import (
	"framework/proto/response"
	"github.com/gin-gonic/gin"
	"service/core/area"
)

// 获取国家
func getCountry(ctx *gin.Context) {
	country := area.GetCountry()
	response.New(ctx).Data(country)
}
