/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-18 23:25:12
 */

package user

import (
	"framework/mysql"
	"log"
	"testing"
)

func TestAutoMigrate(t *testing.T) {
	mysql.AutoMigrate(User{})
}

func TestSave(t *testing.T) {
	user := User{
		Name:     "name",
		Phone:    "13012341234",
		Password: "123456",
		WxOpenId: "wxOpenId",
	}

	mysql.Save(&user)
}

func TestFind(t *testing.T) {
	if one := mysql.FindOne[User](User{
		Name: "name",
	}); one == nil {
		log.Fatalf("one == nil")
	} else {
		log.Println(one.Name)
	}
}
