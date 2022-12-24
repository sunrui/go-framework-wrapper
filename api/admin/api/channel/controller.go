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
	ctx                    *service.Context               // 上下文
	channelRepository      channel.ChannelRepository      // 渠道仓库
	channelGroupRepository channel.ChannelGroupRepository // 渠道组仓库
}

// NewController 创建控制器
func NewController(ctx *service.Context) Controller {
	return Controller{
		ctx:                    ctx,
		channelRepository:      channel.NewChannelRepository(ctx.Mysql),
		channelGroupRepository: channel.NewChannelGroupRepository(ctx.Mysql),
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
