/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-06 17:24:22
 */

package mysql

import (
	"gorm.io/gorm"
)

type Mysql struct {
	*gorm.DB
}

// AutoMigrate 创建表
func (mysql Mysql) AutoMigrate(dst ...any) {
	if err := mysql.DB.AutoMigrate(dst...); err != nil {
		panic(err.Error())
	}
}

// Save 插入
func (mysql Mysql) Save(value any) {
	// 保存新的用户
	if tx := mysql.DB.Save(value); tx.Error != nil {
		panic(tx.Error.Error())
	}
}

// Truncate 清空数据
func (mysql Mysql) Truncate(dst any) {
	mysql.DB.Unscoped().Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&dst)
}
