/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-26 14:53:35
 */

package userDevice

// Type 设备类型
type Type string

const (
	AndroidType Type = "AndroidType" // 安卓
	iOSType     Type = "iOSType"     // 苹果
	WebType     Type = "WebType"     // 网页
)

// LoginType 登录类型
type LoginType string

const (
	PhoneType  LoginType = "phone"      // 手机号
	NameType   LoginType = "name"       // 用户名
	WechatType LoginType = "wechat"     // 微信
	AlipayType LoginType = "alipayType" // 阿里云
)
