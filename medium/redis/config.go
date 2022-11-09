/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-09 22:14:07
 */

package redis

type config struct {
	Host     string `json:"host"`     // 主机
	Port     int    `json:"port"`     // 端口
	Password string `json:"password"` // 密码
	Database int    `json:"database"` // 数据库
}
