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
	AutoMigrate(&User{})
}

func TestMysql_Save(t *testing.T) {
	Truncate(&User{})

	user := User{
		Name: "name",
	}

	Save(&user)
}

func TestFindById(t *testing.T) {
	Truncate(&User{})

	user := User{
		Name: "name",
	}
	Save(&user)

	if one := user.FindById("not found"); one != nil {
		t.Fatalf("one != nil")
	}

	if one := user.FindById(user.Id); one == nil {
		t.Fatalf("one == nil")
	}
}

func TestFindOne(t *testing.T) {
	Truncate(&User{})

	user := User{
		Name: "name",
	}
	Save(&user)

	if one := FindOne[User](User{
		Name: "name-1",
	}); one != nil {
		log.Fatalf("one != nil")
	}

	if one := FindOne[User](User{
		Name: "name",
	}); one == nil {
		log.Fatalf("one == nil")
	}
}

func TestFindPage(t *testing.T) {
	Truncate(&User{})

	for i := 0; i < 10; i++ {
		user := User{
			Name: fmt.Sprintf("name-%d", i),
		}
		Save(&user)
	}

	if ones := FindPage[User](0, 2, "name ASC", User{
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
		if ones := FindPage[User](i, 2, "name ASC", "name LIKE ?", "%name%"); len(ones) != 2 {
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

	if ones := FindPage[User](0, -1, "created_at ASC", nil); len(ones) != 10 {
		for index, value := range ones {
			fmt.Println(index, value)
		}

		t.Fatalf("len(ones) != 10")
	} else {
		for index, value := range ones {
			fmt.Println(index, value)
		}
	}

	if ones := FindPage[User](0, -1, "created_at DESC", "name LIKE ?", "%name%"); len(ones) != 10 {
		for index, value := range ones {
			fmt.Println(index, value)
		}

		t.Fatalf("len(ones) != 10")
	} else {
		for index, value := range ones {
			fmt.Println(index, value)
		}
	}
}

func TestSoftDeleteById(t *testing.T) {
	Truncate(&User{})

	user := User{
		Name: "name",
	}
	Save(&user)

	if r := SoftDeleteById[User](user.Id); r != true {
		log.Fatalf("r != true")
	}

	if r := SoftDeleteById[User](user.Id); r == true {
		log.Fatalf("r == true")
	}
}

func TestDeleteById(t *testing.T) {
	Truncate(&User{})

	user := User{
		Name: "name",
	}
	Save(&user)

	if r := DeleteById[User](user.Id); r != true {
		log.Fatalf("r != true")
	}

	if r := DeleteById[User](user.Id); r == true {
		log.Fatalf("r == true")
	}
}
