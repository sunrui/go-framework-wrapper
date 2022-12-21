/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-12-21 14:44:21
 */

package log

import (
	"framework/app/glog"
	"framework/app/mysql"
)

// LogHttp http 日志
type LogHttp struct {
	mysql.Model            // 通用参数
	Level       glog.Level `json:"level"  gorm:"type:char(5); comment:日志级别"`     // 日志级别
	Ip          string     `json:"ip"  gorm:"type:char(15); comment:ip 地址"`      // ip 地址
	Method      string     `json:"method"  gorm:"type:char(14); comment:请求方式"`   // 请求方式
	Uri         string     `json:"uri"  gorm:"type:varchar(1024); comment:访问地址"` // 访问地址
	Body        *string    `json:"body"  gorm:"type:text; comment:请求体"`          // 请求体
	Response    *string    `json:"response"  gorm:"type:text; comment:返回结果"`     // 返回结果
	UserId      *string    `json:"userId"  gorm:"type:char(12); comment:用户 id"`  // 用户 id
	Elapsed     int64      `json:"elapsed"  gorm:"type:bigint; comment:耗时"`      // 耗时
}

// LogHttpRepository http 仓库
type LogHttpRepository struct {
	Mysql                     *mysql.Mysql // 数据库
	mysql.Repository[LogHttp]              // 通用仓库
}

// NewLogHttpRepository 创建 http 仓库
func NewLogHttpRepository(mysql *mysql.Mysql) LogHttpRepository {
	return LogHttpRepository{
		Mysql: mysql,
	}
}
