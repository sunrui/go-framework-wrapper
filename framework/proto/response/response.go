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
func (responseDef *Response) Ok() {
	responseDef.ctx.JSON(result.Ok.Status, result.Ok)
	responseDef.ctx.Abort()
}

// Data 数据返回对象
func (responseDef *Response) Data(data interface{}) {
	// 判断是否为 Result 对象
	r, ok := data.(result.Result)
	if ok {
		responseDef.ctx.JSON(r.Status, data)
	} else {
		responseDef.ctx.JSON(result.Ok.Status, result.Ok.WithData(data))
	}

	responseDef.ctx.Abort()
}

// IdData 主键返回对象
func (responseDef *Response) IdData(id string) {
	type idData struct {
		Id string `json:"id"`
	}

	responseDef.Data(idData{
		Id: id,
	})
}

// PageData 分页数据返回对象
func (responseDef *Response) PageData(data interface{}, pagination result.Pagination) {
	responseDef.ctx.JSON(result.Ok.Status, result.PageResult{
		Result:     result.Ok.WithData(data),
		Pagination: pagination,
	})
	responseDef.ctx.Abort()
}
