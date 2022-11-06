/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-06 17:53:55
 */

package mysql

import (
	"testing"
)

func TestMysql(t *testing.T) {
	type User struct {
		Model        // 通用参数
		Name  string `json:"name"` // 姓名
	}

	Inst.AutoMigrate(&User{})

	user := User{
		Name: "name",
	}

	Inst.Save(&user)

	one := FindOne[User](User{
		Name: "name1",
	})

	println(one.Name)
}
