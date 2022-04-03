/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022/04/03 00:05:03
 */

package response

import (
	"github.com/gin-gonic/gin"
	"medium-server-go/framework/proto/result"
)

// 结果对象实体
type Response struct {
	ctx *gin.Context
}

// New 新建结果对象
func New(ctx *gin.Context) *Response {
	return &Response{
		ctx: ctx,
	}
}

// Ok 操作成功返回对象
func (response *Response) Ok() {
	response.ctx.JSON(result.Ok.Status, result.Ok)
	response.ctx.Abort()
}

// Data 数据返回对象
func (response *Response) Data(data interface{}) {
	// 判断是否为 Result 对象
	r, ok := data.(result.Result)
	if ok {
		response.ctx.JSON(r.Status, data)
	} else {
		response.ctx.JSON(result.Ok.Status, result.Ok.WithData(data))
	}

	response.ctx.Abort()
}

// IdData 主键返回对象
func (response *Response) IdData(id string) {
	type idData struct {
		Id string `json:"id"`
	}

	response.Data(idData{
		Id: id,
	})
}

// PageData 分页数据返回对象
func (response *Response) PageData(data interface{}, pagination result.Pagination) {
	response.ctx.JSON(result.Ok.Status, result.PageResult{
		Result:     result.Ok.WithData(data),
		Pagination: pagination,
	})
	response.ctx.Abort()
}
