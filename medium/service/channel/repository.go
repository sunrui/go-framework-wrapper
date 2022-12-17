/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-12-17 20:02:16
 */

package channel

import "framework/app/mysql"

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
