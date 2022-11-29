/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-26 20:15:41
 */

package user

// DeviceType 设备类型
type DeviceType string

const (
	AndroidType DeviceType = "ANDROID" // 安卓
	iOSType     DeviceType = "IOS"     // 苹果
	WebType     DeviceType = "WEB"     // 网页
)

// LoginType 登录类型
type LoginType string

const (
	PhoneType  LoginType = "PHONE"  // 手机号
	NameType   LoginType = "NAME"   // 用户名
	WechatType LoginType = "WECHAT" // 微信
	AlipayType LoginType = "ALIPAY" // 阿里云
)

// RoleType 类型
type RoleType string

const (
	AdminType     RoleType = "ADMIN"     // 管理员
	OperationType RoleType = "OPERATION" // 运营
	CockpitType   RoleType = "COCKPIT"   // 驾驶仓
)
