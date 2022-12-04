/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-26 20:14:28
 */

package service

import (
	"medium/service/user"
)

// Mirage 数据库初始化
func Mirage() {
	Context.Mysql.AutoMigrate(&user.User{})
	Context.Mysql.AutoMigrate(&user.Info{})
	Context.Mysql.AutoMigrate(&user.Device{})
	Context.Mysql.AutoMigrate(&user.Role{})
}
