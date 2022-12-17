/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-12-05 03:25:44
 */

package log

import "framework/app/mysql"

// HttpRepository http 仓库
type HttpRepository struct {
	Mysql                     *mysql.Mysql // 数据库
	mysql.Repository[LogHttp]              // 通用仓库
}

// NewHttpRepository 创建 http 仓库
func NewHttpRepository(mysql *mysql.Mysql) HttpRepository {
	return HttpRepository{
		Mysql: mysql,
	}
}
