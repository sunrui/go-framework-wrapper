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
)

// 获取省
func getProvince(_ *gin.Context) result.Result {
	provinces := area.GetProvinces()
	return result.Ok.WithData(provinces)
}
