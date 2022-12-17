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

	ChannelGroup   *ChannelGroup `json:"channelGroup,omitempty" gorm:"foreignKey:ChannelGroupId; comment:渠道组"`
	ChannelGroupId string        `json:"channelGroupId" gorm:"type:char(16); comment:渠道组 id"`

	Enable bool `json:"enable" gorm:"type:tinyint(1); comment:启用"` // 启用
}

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
