/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-12-05 03:25:44
 */

package log

import "framework/app/mysql"

type HttpRepository struct {
	Mysql *mysql.Mysql
	mysql.Repository[Http]
}

func NewHttpRepository(mysql *mysql.Mysql) HttpRepository {
	return HttpRepository{
		Mysql: mysql,
	}
}
