/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/31 20:57:31
 */

package open

import "medium-server-go/framework/db"

// 初始化
func init() {
	var err error

	// 创建表入驻审核
	err = db.Mysql.AutoMigrate(&OpenApproval{})
	if err != nil {
		panic(err.Error())
	}

	// 创建表入驻资料
	err = db.Mysql.AutoMigrate(&OpenSettleIn{})
	if err != nil {
		panic(err.Error())
	}

	// 创建表入驻
	err = db.Mysql.AutoMigrate(&Open{})
	if err != nil {
		panic(err.Error())
	}
}

// 获取指定用户下所有入驻
func GetOpens(userId string) []Open {
	var open []Open

	query := db.Mysql.Where(Open{
		UserId: userId,
	}).Find(&open)

	if query.Error != nil {
		panic(query.Error.Error())
	}

	return open
}
