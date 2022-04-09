/*
 * Copyright (c) $today.year honeysense.com All rights reserved.
 * Author: $author
 * Date: $today.format("yyyy-MM-dd HH:mm:ss")
 */

package template

import "framework/db"

// Template 模板
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
