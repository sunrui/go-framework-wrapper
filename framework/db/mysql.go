/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/07 03:02:07
 */

package db

import (
	"fmt"
	"framework/config"
	"framework/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

// Mysql 数据库访问对象
var Mysql *gorm.DB

// Model 数据库通用对象
type Model struct {
	Id        string     `json:"id" gorm:"primaryKey;type:varchar(12);comment:主键 id"` // 主键 id
	CreatedAt time.Time  `json:"created_at" gorm:"autoCreateTime:milli;comment:创建时间"` // 创建时间
	UpdatedAt time.Time  `json:"updated_at" gorm:"autoUpdateTime:milli;comment:更新时间"` // 更新时间
	DeletedAt *time.Time `json:"deleted_at,omitempty" gorm:"comment:删除时间"`            // 删除时间
}

// BeforeCreate 创建对象前回调
func (model *Model) BeforeCreate(*gorm.DB) (err error) {
	model.Id = utils.CreateNanoid()
	return nil
}

// 初始化
func init() {
	var err error

	// 数据库连接
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		config.Get().Mysql.User,
		config.Get().Mysql.Password,
		config.Get().Mysql.Host,
		config.Get().Mysql.Port,
		config.Get().Mysql.Database)

	if Mysql, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_", // 表名前缀
			SingularTable: true, // 使用单数表名
		},
	}); err != nil {
		panic(err.Error())
	}
}
