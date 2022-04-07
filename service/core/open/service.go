/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/31 20:57:31
 */

package open

import (
	"framework/db"
)

// GetOpen 获取指定用户下所有入驻
func GetOpen(userId string) []Open {
	var open []Open

	query := db.Mysql.Where(Open{
		UserId: userId,
	}).Find(&open)

	if query.Error != nil {
		panic(query.Error.Error())
	}

	return open
}
