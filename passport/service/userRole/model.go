/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-19 00:36:52
 */

package userRole

import "framework/mysql"

// UserRole 用户角色
type UserRole struct {
	mysql.Model[UserRole]
	UserId string `json:"userId" gorm:"type:char(12);comment:用户 id"` // 用户 id
	Type   Type   `json:"type" gorm:"type:varchar, comment:类型"`      // 类型
}

func init() {
	mysql.AutoMigrate(&UserRole{})
}
