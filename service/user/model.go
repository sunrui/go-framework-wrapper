/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/21 00:23:21
 */

package user

import "medium-server-go/framework/db"

// User 用户
type User struct {
	db.Model         // 通用参数
	Phone     string `json:"phone" gorm:"uniqueIndex, not null; comment:手机号"` // 手机号
	Ip        string `json:"ip" gorm:"comment:ip 地址"`                         // ip 地址
	UserAgent string `json:"userAgent" gorm:"comment:用户 ua"`                  // 用户 ua
}

// Save 存储
func (user *User) Save() {
	db.Mysql.Save(user)
}

// 初始化
func init() {
	// 创建表用户
	err := db.Mysql.AutoMigrate(&User{})
	if err != nil {
		panic(err.Error())
	}
}
