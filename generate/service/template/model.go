/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-04-07 16:12:55
 */

package template

import (
	"medium-server-go/framework/db"
)

// Template ${Template.name}
type Template struct {
	db.Model        // 通用参数
	UserId   string `json:"userId"` // 用户 id
	Name     string `json:"name"`   // 名称
}

// 初始化
func init() {
	var err error

	// 创建表入驻
	err = db.Mysql.AutoMigrate(&Template{})
	if err != nil {
		panic(err.Error())
	}
}

// Save 保存
func (template *Template) Save() {
	db.Mysql.Save(template)
}
