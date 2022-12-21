/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-12-17 18:48:17
 */

package channel

import "framework/app/mysql"

// Channel 渠道
type Channel struct {
	mysql.Model
	UserId string `json:"userId" gorm:"type:char(16); comment:用户 id"`       // 用户 id
	Name   string `json:"name" gorm:"type:varchar(32); unique; comment:名称"` // 名称

	ChannelGroup   *ChannelGroup `json:"channelGroup,omitempty" gorm:"foreignKey:ChannelGroupId; comment:渠道组"` // 渠道组
	ChannelGroupId string        `json:"channelGroupId" gorm:"type:char(16); comment:渠道组 id"`                  // 渠道组 id

	Enable bool `json:"enable" gorm:"type:tinyint(1); comment:启用"` // 启用
}

// ChannelRepository 渠道仓库
type ChannelRepository struct {
	Mysql                     *mysql.Mysql // 数据库
	mysql.Repository[Channel]              // 通用仓库
}

// NewChannelRepository 创建渠道仓库
func NewChannelRepository(mysql *mysql.Mysql) ChannelRepository {
	return ChannelRepository{
		Mysql: mysql,
	}
}
