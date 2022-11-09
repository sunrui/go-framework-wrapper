/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-06 23:34:42
 */

package user

import "medium/mysql"

type User struct {
	mysql.Model[User]
	Name string `json:"name"`
}

func FindByName(name string) *User {
	return nil
}
