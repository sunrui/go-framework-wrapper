/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-18 22:33:51
 */

package user

import "framework/mysql"

type User struct {
	mysql.Model[User]
	Name         string `json:"name" gorm:"index;unique;comment:用户名"`  // 用户名
	Phone        string `json:"phone" gorm:"index;unique;comment:手机号"` // 手机号
	Password     string `json:"password" gorm:"comment:密码"`            // 密码
	WxOpenId     string `json:"wxOpenId" gorm:"comment:微信 openId"`     // 微信 openId
	AlipayOpenId string `json:"aliOpenId" gorm:"comment:支付宝 openId"`   // 支付宝 openId
}
