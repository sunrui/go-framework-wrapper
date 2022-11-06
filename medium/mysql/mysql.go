/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-06 17:24:22
 */

package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Mysql struct {
	*gorm.DB
}

// NewMysql 创建对象
func NewMysql(conf conf) *Mysql {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.User,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.Database)

	if db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_", // 表名前缀
			SingularTable: true, // 使用单数表名
		},
	}); err != nil {
		panic(err.Error())
	} else {
		return &Mysql{
			DB: db,
		}
	}
}

// AutoMigrate 创建表
func (mysql Mysql) AutoMigrate(dst ...any) {
	if err := mysql.DB.AutoMigrate(dst...); err != nil {
		panic(err.Error())
	}
}

// 插入
func (mysql Mysql) Save(value any) {
	// 保存新的用户
	if tx := mysql.DB.Save(value); tx.Error != nil {
		panic(tx.Error.Error())
	}
}
