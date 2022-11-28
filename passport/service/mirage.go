/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-26 20:14:28
 */

package service

import (
	"framework/context"
	"passport/service/user"
)

// Mirage 数据库初始化
func Mirage() {
	context.Mysql.AutoMigrate(&user.User{})
	context.Mysql.AutoMigrate(&user.Info{})
	context.Mysql.AutoMigrate(&user.Device{})
	context.Mysql.AutoMigrate(&user.Role{})
}
