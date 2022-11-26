/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-26 13:12:34
 */

package user

import (
	"framework/mysql"
	"golang.org/x/crypto/bcrypt"
)

func IsValidateNameAndPassword(name string, password string) bool {
	if user := mysql.FindOne[User](User{
		Name: name,
	}); user == nil {
		return false
	} else {
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err == nil {
			return true
		}
	}

	return false
}

func IsValidatePhoneAndPassword(phone string, password string) bool {
	if user := mysql.FindOne[User](User{
		Phone: phone,
	}); user == nil {
		return false
	} else {
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err == nil {
			return true
		}
	}

	return false
}
