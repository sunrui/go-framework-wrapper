/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-12-05 03:12:48
 */

package log

import (
	"framework/app/glog"
	"framework/app/mysql"
)

type Http struct {
	mysql.Model            // 通用参数
	Level       glog.Level `json:"level"  gorm:"comment:日志级别"`    // 日志级别
	Ip          string     `json:"ip"  gorm:"comment:ip 地址"`      // ip 地址
	Method      string     `json:"method"  gorm:"comment:请求方式"`   // 请求方式
	Uri         string     `json:"uri"  gorm:"comment:访问地址"`      // 访问地址
	Body        *string    `json:"body"  gorm:"comment:请求体"`      // 请求体
	Response    *string    `json:"response"  gorm:"comment:返回结果"` // 返回结果
	UserId      *string    `json:"userId"  gorm:"comment:用户 id"`  // 用户 id
	Elapsed     int64      `json:"elapsed"  gorm:"comment:耗时"`    // 耗时
}

func (Http) TableName() string {
	return "t_log_http"
}
