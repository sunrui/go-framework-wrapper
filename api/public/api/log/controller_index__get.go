/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-12-05 04:21:56
 */

package log

import (
	"framework/app/glog"
	"framework/app/result"
	"framework/app/server"
	"github.com/gin-gonic/gin"
	"medium/service/log"
)

type getIndexReq struct {
	result.Page             // 分页
	Level       *glog.Level `json:"level" form:"level" validate:"omitempty,oneof=Debug Info Warn Error"` // 日志级别
}

type getIndexRes struct {
	Http []log.LogHttp
}

// @Summary 列表
// @Tags    日志
// @Produce json
// @Success 200 {object} result.Result{data=log.LogHttp}
// @Router  /public/log [get]
func (controller Controller) getIndex(ctx *gin.Context) result.Result {
	var req getIndexReq

	server.ValidateParameter(ctx, &req)

	return result.Ok.WithData(req)

	// var query log.LogHttp
	// if req.Level != nil {
	//	query = log.LogHttp{
	//		Level: *req.Level,
	//	}
	// } else {
	//	query = log.LogHttp{}
	// }

	// var res []log.LogHttp
	// res = controller.LogHttpRepository.FindPage(req.Page, req.Page.PageSize, "ASC", query)
	// return result.Ok.WithData(res)
}
