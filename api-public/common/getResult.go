/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022/04/03 13:44:03
 */

package common

import (
	"framework/proto/response"
	"framework/proto/result"
	"github.com/gin-gonic/gin"
)

// @Summary  结果列表
// @Tags     通用
// @Produce  json
// @Success  200  {object}  result.Result{data=result.Result}  true  {"status":"Ok","description":"成功"}
// @Router   /public/common/result [get]
func getResult(ctx *gin.Context) {
	response.New(ctx).Data(result.All())
}
