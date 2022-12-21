/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-12-21 15:32:12
 */

package user

import "framework/app/mysql"

// LoginType 登录类型
type LoginType string

const (
	PhoneType  LoginType = "PHONE"  // 手机号
	NameType   LoginType = "NAME"   // 用户名
	WechatType LoginType = "WECHAT" // 微信
	AlipayType LoginType = "ALIPAY" // 阿里云
)

// DeviceType 设备类型
type DeviceType string

const (
	AndroidType DeviceType = "ANDROID" // 安卓
	iOSType     DeviceType = "IOS"     // 苹果
	WebType     DeviceType = "WEB"     // 网页
)

type UserDevice struct {
	mysql.Model
	UserId      string     `json:"userId"`      // 用户 id
	Type        DeviceType `json:"type"`        // 类型
	Ip          string     `json:"ip"`          // ip
	PackageName string     `json:"packageName"` // 包名
	AppVersion  string     `json:"appVersion"`  // 软件版本
	JwtToken    string     `json:"jwtToken"`    // jwt 令牌
	LoginType   LoginType  `json:"loginType"`   // 登录类型
	UserAgent   string     `json:"userAgent"`   // 用户代理
}
