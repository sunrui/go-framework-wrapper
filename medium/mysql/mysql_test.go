/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-06 17:53:55
 */

package mysql

import (
	"fmt"
	"log"
	"testing"
)

type User struct {
	Model[User]        // 通用参数
	Name        string `json:"name"` // 姓名
}

func TestMysql_AutoMigrate(t *testing.T) {
	Inst.AutoMigrate(&User{})
}

func TestMysql_Save(t *testing.T) {
	Inst.Truncate(&User{})

	user := User{
		Name: "name",
	}

	Inst.Save(&user)
}

func TestFindById(t *testing.T) {
	Inst.Truncate(&User{})

	user := User{
		Name: "name",
	}
	Inst.Save(&user)

	if one := user.FindById("not found"); one != nil {
		t.Fatalf("one != nil")
	}

	if one := user.FindById(user.Id); one == nil {
		t.Fatalf("one == nil")
	}
}

func TestFindOne(t *testing.T) {
	Inst.Truncate(&User{})

	user := User{
		Name: "name",
	}
	Inst.Save(&user)

	if one := user.FindOne(User{
		Name: "name-1",
	}); one != nil {
		log.Fatalf("one != nil")
	}

	if one := user.FindOne(User{
		Name: "name",
	}); one == nil {
		log.Fatalf("one == nil")
	}
}

func TestFindPage(t *testing.T) {
	Inst.Truncate(&User{})

	for i := 0; i < 10; i++ {
		user := User{
			Name: fmt.Sprintf("name-%d", i),
		}
		Inst.Save(&user)
	}

	user := User{}
	if ones := user.FindPage(0, 2, User{
		Name: "hello",
	}); len(ones) != 0 {
		for index, value := range ones {
			fmt.Println(index, value)
		}

		t.Fatalf("len(ones) != 0")
	} else {
		for index, value := range ones {
			fmt.Println(index, value)
		}
	}

	for i := 0; i < 5; i++ {
		// find like name
		if ones := user.FindPage(i, 2, "name LIKE ?", "%name%"); len(ones) != 2 {
			for index, value := range ones {
				fmt.Println(index, value)
			}

			t.Fatalf("len(ones) != 2")
		} else {
			for index, value := range ones {
				fmt.Println(index, value)
			}
		}
	}
}

func TestSoftDeleteById(t *testing.T) {
	Inst.Truncate(&User{})

	user := User{
		Name: "name",
	}
	Inst.Save(&user)

	if r := user.SoftDeleteById(user.Id); r != true {
		log.Fatalf("r != true")
	}

	if r := user.SoftDeleteById(user.Id); r == true {
		log.Fatalf("r == true")
	}
}

func TestDeleteById(t *testing.T) {
	Inst.Truncate(&User{})

	user := User{
		Name: "name",
	}
	Inst.Save(&user)

	if r := user.DeleteById(user.Id); r != true {
		log.Fatalf("r != true")
	}

	if r := user.DeleteById(user.Id); r == true {
		log.Fatalf("r == true")
	}
}
