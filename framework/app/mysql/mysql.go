/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-06 17:24:22
 */

package mysql

import (
	"errors"
	"fmt"
	"framework/app/glog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
		Logger: getLogger(glog),
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

// FindById 根据 id 查找
func FindById[T any](mysql Mysql, id string) *T {
	var dst T

	if db := mysql.DB.Where("id = ?", id).Find(&dst); db.Error != nil {
		panic(db.Error.Error())
	} else if db.RowsAffected == 1 {
		return &dst
	} else {
		return nil
	}
}

// FindOne 根据条件查找一个
func FindOne[T any](mysql *Mysql, query interface{}, args ...interface{}) *T {
	var dst []T

	if db := mysql.DB.Limit(2).Where(query, args).Find(&dst); db.Error != nil {
		panic(db.Error.Error())
	} else if db.RowsAffected > 1 {
		panic(errors.New(fmt.Sprintf("find %d record", db.RowsAffected)))
	} else if db.RowsAffected == 1 {
		return &dst[0]
	} else {
		return nil
	}
}

// FindPage 根据条件查找分页一个或多个
func FindPage[T any](mysql *Mysql, page int, pageSize int, order string, query interface{}, args ...interface{}) []T {
	var dst []T

	var db *gorm.DB

	if query != nil {
		db = mysql.DB.Order(order).Offset(page*pageSize).Limit(pageSize).Offset(page*pageSize).Where(query, args).Find(&dst)
	} else {
		db = mysql.DB.Order(order).Offset(page * pageSize).Limit(pageSize).Offset(page * pageSize).Find(&dst)
	}

	if db.Error != nil {
		panic(db.Error.Error())
	} else {
		return dst
	}
}

// SoftDeleteById 根据 id 删除
func SoftDeleteById[T any](mysql *Mysql, id string) bool {
	var dst T

	if r := mysql.DB.Where("id = ?", id).Delete(&dst); r.Error != nil {
		panic(r.Error.Error())
	} else if r.RowsAffected >= 1 {
		return true
	} else {
		return false
	}
}

// DeleteById 根据 id 删除
func DeleteById[T any](mysql *Mysql, id string) bool {
	var dst T

	if r := mysql.DB.Unscoped().Where("id = ?", id).Delete(&dst); r.Error != nil {
		panic(r.Error.Error())
	} else if r.RowsAffected >= 1 {
		return true
	} else {
		return false
	}
}