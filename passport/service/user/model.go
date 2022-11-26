/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-18 22:33:51
 */

package user

import (
	"framework/mysql"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	mysql.Model
	Name         string `json:"name" gorm:"index;unique;comment:用户名"`  // 用户名
	Phone        string `json:"phone" gorm:"index;unique;comment:手机号"` // 手机号
	Password     string `json:"password" gorm:"binary(60);comment:密码"` // 密码
	WxOpenId     string `json:"wxOpenId" gorm:"comment:微信 openId"`     // 微信 openId
	AlipayOpenId string `json:"aliOpenId" gorm:"comment:支付宝 openId"`   // 支付宝 openId
}

// BeforeSave 更改密码
func (user *User) BeforeSave(tx *gorm.DB) error {
	if pw, err := bcrypt.GenerateFromPassword([]byte(user.Password), 0); err == nil {
		tx.Statement.SetColumn("password", pw)
	} else {
		return err
	}

	return nil
}

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

type UserInfo struct {
	mysql.Model
}

// UserRole 用户角色
type UserRole struct {
	mysql.Model
	UserId string `json:"userId" gorm:"type:char(12);comment:用户 id"` // 用户 id
	//Type   RoleType `json:"type" gorm:"type:varchar(32), comment:类型"`  // 类型
}
