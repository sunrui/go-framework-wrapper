/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-12-21 14:43:08
 */

package channel

import "framework/app/mysql"

// ChannelGroupModel 渠道组 - 模型
type ChannelGroupModel struct {
	Name   string `json:"name" gorm:"type:varchar(32); unique; comment:名称" validate:"required, max=32"` // 名称
	AreaId int    `json:"areaId" gorm:"type:int; comment:区 id" validate:"required"`                     // 区 id
}

// ChannelGroup 渠道组
type ChannelGroup struct {
	mysql.Model
	UserId string `json:"userId" gorm:"type:char(16); comment:用户 id" form:"level" ` // 用户 id
	ChannelGroupModel
	Enable bool `json:"enable" gorm:"type:tinyint(1); comment:启用"` // 启用
}

// ChannelGroupRepository 渠道组仓库
type ChannelGroupRepository struct {
	Mysql                          *mysql.Mysql // 数据库
	mysql.Repository[ChannelGroup]              // 通用仓库
}

func NewChannelGroupRepository(mysql *mysql.Mysql) ChannelGroupRepository {
	return ChannelGroupRepository{
		Mysql: mysql,
	}
}
