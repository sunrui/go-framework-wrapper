/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-04-07 16:12:11
 */

package template

import (
	"github.com/gin-gonic/gin"
	"medium-server-go/framework/app"
	"medium-server-go/framework/proto/request"
	"medium-server-go/framework/proto/token"
)

// @Summary  获取所有验证码
// @Tags      演示
// @Accept    json
// @Produce   json
// @Param    page      query     int                            false  "分页"
// @Param    pageSize  query     int                            false  "分页大小"
// @Success  200       {object}  result.PageResult{data=[]Sms}  true
// @Router   /sms [get]
func getOne(ctx *gin.Context) {
	var req request.PageRequest

	// 获取当前 userId
	userId := token.GetUserId(ctx)

	// 较验参数
	app.ValidateParameter(ctx, &req)

	println(userId)

	//response.New(ctx).PageData(all(), result.Pagination{
	//	Page:      req.Page,
	//	PageSize:  req.PageSize,
	//	TotalPage: 10,
	//	TotalSize: 100,
	//})
}
