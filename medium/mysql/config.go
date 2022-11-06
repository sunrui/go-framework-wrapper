/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-06 19:35:52
 */

package mysql

type conf struct {
	User     string `json:"user"`     // 用户名
	Password string `json:"password"` // 密码
	Host     string `json:"host"`     // 主机
	Port     int    `json:"port"`     // 端口
	Database string `json:"database"` // 数据库
}
