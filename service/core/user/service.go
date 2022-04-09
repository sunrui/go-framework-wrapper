/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/21 00:45:21
 */

package user

import (
	"errors"
	"framework/db"
	"gorm.io/gorm"
)

// FindById 根据 id 获取用户
func FindById(id string) *User {
	var user User

	query := db.Mysql.Find(&user, id)
	if errors.Is(query.Error, gorm.ErrRecordNotFound) {
		return nil
	}

	return &user
}

// FindByPhone 根据 id 获取用户
func FindByPhone(phone string) *User {
	var user User

	query := db.Mysql.First(&user, "phone = ?", phone)
	if errors.Is(query.Error, gorm.ErrRecordNotFound) {
		return nil
	}

	return &user
}