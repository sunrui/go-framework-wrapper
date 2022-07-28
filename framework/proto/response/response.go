/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022/04/03 00:05:03
 */

package response

import (
	"framework/proto/result"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response 结果对象实体
type Response struct {
	ctx *gin.Context
}

// New 新建结果对象
func New(ctx *gin.Context) *Response {
	return &Response{
		ctx: ctx,
	}
}

// Data 数据返回对象
func (response *Response) Data(data any) {
	response.ctx.JSON(http.StatusOK, data)
}

// Ok 操作成功返回对象
func (response *Response) Ok() {
	response.Data(result.Ok)
}

// Id 主键返回对象
func (response *Response) Id(id string) {
	response.Data(result.Ok.WithIdData(id))
}

// Result 数据返回对象
func (response *Response) Result(result result.Result) {
	response.Data(result)
}

// PageResult 数据返回对象
func (response *Response) PageResult(pageResult result.PageResult) {
	response.Data(pageResult)
}
