/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-26 13:15:54
 */

package user

import (
	"framework/mysql"
	"testing"
)

func TestIsValidateNameAndPassword(t *testing.T) {
	user := User{
		Name:     "name",
		Password: "123456",
	}

	if one := mysql.FindOne[User](user); one == nil {
		mysql.Save(&user)
	} else {
		mysql.DeleteById[User](one.Id)
		mysql.Save(&user)
	}

	if ok := IsValidateNameAndPassword(user.Name, "12345"); ok {
		t.Error("password is error")
	}

	if ok := IsValidateNameAndPassword(user.Name, user.Password); !ok {
		t.Error("password is correct")
	}
}

func TestIsValidatePhoneAndPassword(t *testing.T) {
	user := User{
		Phone:    "13012341234",
		Password: "123456",
	}

	if one := mysql.FindOne[User](user); one == nil {
		mysql.Save(&user)
	} else {
		mysql.DeleteById[User](one.Id)
		mysql.Save(&user)
	}

	if ok := IsValidatePhoneAndPassword(user.Phone, "12345"); ok {
		t.Error("password is error")
	}

	if ok := IsValidatePhoneAndPassword(user.Phone, user.Password); !ok {
		t.Error("password is correct")
	}
}
