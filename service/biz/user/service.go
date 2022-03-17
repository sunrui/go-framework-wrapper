/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/21 00:45:21
 */

package user

import (
	"errors"
	"gorm.io/gorm"
	"medium-server-go/common/db"
)

// 初始化
func init() {
	// 创建表
	err := db.Mysql.AutoMigrate(&User{})
	if err != nil {
		panic(err.Error())
	}
}

// 存储数据
func SaveUser(user *User) {
	db.Mysql.Save(user)
}

// 根据 id 获取用户
func FindByPhone(phone string) *User {
	var user User

	query := db.Mysql.First(&user, "phone = ?", phone)
	if errors.Is(query.Error, gorm.ErrRecordNotFound) {
		return nil
	}

	return &user
}

// 根据 id 获取用户
func FindById(id string) *User {
	var user User

	query := db.Mysql.Find(&user, id)
	if errors.Is(query.Error, gorm.ErrRecordNotFound) {
		return nil
	}

	return &user
}
