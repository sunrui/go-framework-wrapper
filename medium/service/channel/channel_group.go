/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-12-21 14:43:08
 */

package channel

import "framework/app/mysql"

// ChannelGroup 渠道组
type ChannelGroup struct {
	mysql.Model
	UserId       string `json:"userId" gorm:"type:char(16); comment:用户 id"`        // 用户 id
	Name         string `json:"name" gorm:"type:varchar(32); unique; comment:名称"`  // 名称
	ProvinceId   int    `json:"provinceId" gorm:"type:int; comment:省 id"`          // 省 id
	ProvinceName string `json:"provinceName" gorm:"type:varchar(32); comment:省名称"` // 省名称
	CityId       int    `json:"cityId" gorm:"type:int; comment:市 id"`              // 市 id
	CityName     string `json:"cityName" gorm:"type:varchar(32); comment:市名称"`     // 市名称
	AreaId       int    `json:"areaId" gorm:"type:int; comment:区 id"`              // 区 id
	AreaName     string `json:"areaName" gorm:"type:varchar(32); comment:区名称"`     // 区名称
	Enable       bool   `json:"enable" gorm:"type:tinyint(1); comment:启用"`         // 启用
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
