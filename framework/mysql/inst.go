/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-07 00:26:26
 */

package mysql

import (
	"config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// Mysql 数据库
type Mysql struct {
	*gorm.DB
}

// Inst 实例
var Inst *Mysql

func init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		config.Inst().Mysql.User,
		config.Inst().Mysql.Password,
		config.Inst().Mysql.Host,
		config.Inst().Mysql.Port,
		config.Inst().Mysql.Database)

	if db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: func() logger.Interface {
			if config.IsDev() {
				return logger.Default.LogMode(logger.Info)
			} else {
				return logger.Default.LogMode(logger.Warn)
			}
		}(),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_", // 表名前缀
			SingularTable: true, // 使用单数表名
		},
	}); err == nil {
		// 配置连接池
		sqlDb, _ := db.DB()
		sqlDb.SetMaxOpenConns(config.Inst().Mysql.MaxOpenConns)
		sqlDb.SetMaxIdleConns(config.Inst().Mysql.MaxIdleConns)

		Inst = &Mysql{
			DB: db,
		}
	} else {
		panic(err.Error())
	}
}
