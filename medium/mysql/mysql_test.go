/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-06 17:53:55
 */

package mysql

import (
	"testing"
)

type User struct {
	Model        // 通用参数
	Name  string `json:"name"` // 姓名
}

func TestMysql_AutoMigrate(t *testing.T) {
	Inst.AutoMigrate(&User{})
}

func TestMysql_Save(t *testing.T) {
	user := User{
		Name: "name",
	}
	Inst.Save(&user)
}

func TestFindOne(t *testing.T) {
	one := FindOne[User](User{
		Name: "name-1",
	})

	if one != nil {
		t.Fatalf("one is not nil")
	}
}
