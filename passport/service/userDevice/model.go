/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-26 14:56:29
 */

package userDevice

import "framework/mysql"

type UserDevice struct {
	mysql.Model[UserDevice]
	UserId      string    `json:"userId"`      // 用户 id
	Type        Type      `json:"type"`        // 类型
	Ip          string    `json:"ip"`          // ip
	PackageName string    `json:"packageName"` // 包名
	AppVersion  string    `json:"appVersion"`  // 软件版本
	JwtToken    string    `json:"jwtToken"`    // jwt 令牌
	LoginType   LoginType `json:"loginType"`   // 登录类型
	UserAgent   string    `json:"userAgent"`   // 用户代理
}

func init() {
	mysql.AutoMigrate(&UserDevice{})
}
