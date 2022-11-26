/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-26 20:15:41
 */

package user

// DeviceType 设备类型
type DeviceType string

const (
	AndroidType DeviceType = "AndroidType" // 安卓
	iOSType     DeviceType = "iOSType"     // 苹果
	WebType     DeviceType = "WebType"     // 网页
)

// LoginType 登录类型
type LoginType string

const (
	PhoneType  LoginType = "phone"      // 手机号
	NameType   LoginType = "name"       // 用户名
	WechatType LoginType = "wechat"     // 微信
	AlipayType LoginType = "alipayType" // 阿里云
)

// RoleType 类型
type RoleType string

const (
	AdminType     RoleType = "Admin"     // 管理员
	OperationType RoleType = "Operation" // 运营
	CockpitType   RoleType = "Cockpit"   // 驾驶仓
)
