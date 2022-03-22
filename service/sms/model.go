/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/16 05:44:16
 */

package sms

import (
	"medium-server-go/framework/db"
)

// 验证码
type Code struct {
	db.Model           // 通用参数
	Phone     string   `json:"phone" gorm:"comment:手机号"`       // 手机号
	CodeType  CodeType `json:"codeType" gorm:"comment:短信类型"`   // 短信类型
	Code      string   `json:"code" gorm:"comment:验证码"`        // 验证码
	Ip        string   `json:"ip" gorm:"comment:ip 地址"`        // ip 地址
	UserAgent string   `json:"userAgent" gorm:"comment:用户 ua"` // 用户 ua
	Success   bool     `json:"success" gorm:"comment:是否发送成功"`  // 是否发送成功
	Comment   string   `json:"comment" gorm:"comment:备注"`      // 备注
}

// 存储
func (code *Code) Save() {
	db.Mysql.Save(code)
}

// 初始化
func init() {
	// 创建表验证码
	err := db.Mysql.AutoMigrate(&Code{})
	if err != nil {
		panic(err.Error())
	}
}
