/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-04-07 10:58:35
 */

package service

import "medium-server-go/framework/db"

// Generate 对象
type Generate struct {
	db.Model        // 通用参数
	UserId   string `json:"userId"` // 用户 id
}

// 初始化
func init() {
	var err error

	// 创建表入驻
	err = db.Mysql.AutoMigrate(&Generate{})
	if err != nil {
		panic(err.Error())
	}
}

// Save 存储
func (generate *Generate) Save() {
	db.Mysql.Save(generate)
}
