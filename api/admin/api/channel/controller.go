/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-12-17 21:40:47
 */

package channel

import (
	"framework/app/server"
	"medium/service"
	"medium/service/channel"
	"net/http"
)

// Controller 控制器
type Controller struct {
	Ctx                    *service.Context               // 上下文
	ChannelRepository      channel.ChannelRepository      // 渠道仓库
	ChannelGroupRepository channel.ChannelGroupRepository // 渠道组仓库
}

// NewController 创建控制器
func NewController(ctx *service.Context) Controller {
	return Controller{
		Ctx:                    ctx,
		ChannelRepository:      channel.NewChannelRepository(ctx.Mysql),
		ChannelGroupRepository: channel.NewChannelGroupRepository(ctx.Mysql),
	}
}

// GetRouter 获取路由
func (controller Controller) GetRouter() server.RouterGroup {
	return server.RouterGroup{
		GroupName:  "/channel",
		Middleware: nil,
		Routers: []server.Router{
			{
				HttpMethod:   http.MethodPost,
				RelativePath: "group",
				RouterFunc:   controller.postChannelGroup,
			},
			{
				HttpMethod:   http.MethodDelete,
				RelativePath: "group/:channelGroupId",
				RouterFunc:   controller.deleteChannelGroup,
			},
			{
				HttpMethod:   http.MethodGet,
				RelativePath: "group",
				RouterFunc:   controller.getChannelGroup,
			},
			{
				HttpMethod:   http.MethodPost,
				RelativePath: "/",
				RouterFunc:   controller.postChannel,
			},
			{
				HttpMethod:   http.MethodDelete,
				RelativePath: "/:channelId",
				RouterFunc:   controller.deleteChannel,
			},
			{
				HttpMethod:   http.MethodGet,
				RelativePath: "/",
				RouterFunc:   controller.getChannel,
			},
		},
	}
}
