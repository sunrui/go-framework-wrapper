/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-18 23:00:30
 */

package userInfo

import "framework/mysql"

type UserInfo struct {
	mysql.Model[UserInfo]
}

func init() {
	mysql.AutoMigrate(&UserInfo{})
}
