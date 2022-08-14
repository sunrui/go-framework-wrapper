/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/21 00:45:21
 */

package user

import (
	"framework/db"
)

// FindById 根据 id 获取用户
func FindById(id string) *User {
	var user User

	if tx := db.Mysql.Find(&user, id); tx.Error != nil {
		panic(tx.Error.Error())
	} else if tx.RowsAffected == 1 {
		return &user
	} else {
		return nil
	}

	return &user
}

// FindByPhone 根据 id 获取用户
func FindByPhone(phone string) *User {
	var user User

	if tx := db.Mysql.Find(&user, "phone = ?", phone); tx.Error != nil {
		panic(tx.Error.Error())
	} else if tx.RowsAffected == 1 {
		return &user
	} else {
		return nil
	}
}
