/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-06 17:24:22
 */

package mysql

import (
	"fmt"
	"framework/app/glog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// Config 配置
type Config struct {
	User         string `json:"user"`         // 用户名
	Password     string `json:"password"`     // 密码
	Host         string `json:"host"`         // 主机
	Port         int    `json:"port"`         // 端口
	Database     string `json:"database"`     // 数据库
	MaxOpenConns int    `json:"maxOpenConns"` // 最大打开连接
	MaxIdleConns int    `json:"maxIdleConns"` // 最大空闲连接
}

// Mysql 数据库
type Mysql struct {
	*gorm.DB
}

// New 创建
func New(config Config, glog *glog.GLog) (*Mysql, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Database)

	if db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: func() logger.Interface {
			if glog != nil {
				return getLogger(glog)
			} else {
				return nil
			}
		}(),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_", // 表名前缀
			SingularTable: true, // 使用单数表名
		},
	}); err == nil {
		// 配置连接池
		sqlDb, _ := db.DB()
		sqlDb.SetMaxOpenConns(config.MaxOpenConns)
		sqlDb.SetMaxIdleConns(config.MaxIdleConns)

		return &Mysql{
			DB: db,
		}, nil
	} else {
		return nil, err
	}
}

// AutoMigrate 创建表
func (mysql Mysql) AutoMigrate(dst ...any) {
	if err := mysql.DB.AutoMigrate(dst...); err != nil {
		panic(err.Error())
	}
}

// Save 插入
func (mysql Mysql) Save(value any) {
	if tx := mysql.DB.Save(value); tx.Error != nil {
		panic(tx.Error.Error())
	}
}

// Truncate 清空数据
func (mysql Mysql) Truncate(dst any) {
	mysql.DB.Unscoped().Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&dst)
}
