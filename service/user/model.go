/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/21 00:23:21
 */

package user

import "medium-server-go/framework/db"

type User struct {
	db.Model         // 通用参数
	Phone     string `json:"phone" gorm:"uniqueIndex, not null; comment:手机号"` // 手机号
	Ip        string `json:"ip" gorm:"comment:ip 地址"`                         // ip 地址
	UserAgent string `json:"userAgent" gorm:"comment:用户 ua"`                  // 用户 ua
}
