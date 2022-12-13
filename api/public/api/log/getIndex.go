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
	"framework/app/server/request"
	"github.com/gin-gonic/gin"
	"medium/service/log"
)

type getIndexReq struct {
	request.Page
	Level *glog.Level
}

type getIndexRes struct {
	Http []log.Http
}

// @Summary 列表
// @Tags    日志
// @Produce json
// @Success 200 {object} result.Result{data=log.Http}
// @Router  /public/log [get]
func (controller Controller) getIndex(ctx *gin.Context) *result.Result {
	var req getIndexReq

	server.ValidateParameter(ctx, &req)

	var query log.Http
	if req.Level != nil {
		query = log.Http{
			Level: *req.Level,
		}
	} else {
		query = log.Http{}
	}

	var res []log.Http
	res = controller.HttpRepository.FindPage(req.Page.Page, req.PageSize, "ASC", query)
	return result.Ok.WithData(res)
}
