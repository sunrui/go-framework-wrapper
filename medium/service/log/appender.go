/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-12-05 03:29:36
 */

package log

import (
	"framework/app/glog"
	"framework/app/mysql"
	"framework/app/util"
)

// Appender 附加者
type Appender struct {
	LogHttpRepository LogHttpRepository
	Mysql             *mysql.Mysql
}

// NewAppender 创建附加者
func NewAppender(mysql *mysql.Mysql) Appender {
	return Appender{
		Mysql:             mysql,
		LogHttpRepository: NewLogHttpRepository(mysql),
	}
}

// Print 打印
func (appender Appender) Print(_ glog.Level, _ string) {
	panic("cannot invoke this method")
}

// PrintHttp 打印消息
func (appender Appender) PrintHttp(level glog.Level, http glog.Http) {
	appender.Mysql.Save(&LogHttp{
		Level:  level,
		Ip:     http.Result.Request.Ip,
		Method: http.Result.Request.Method,
		Uri:    http.Result.Request.Uri,
		Body: func() *string {
			if http.Result.Request.Body != nil {
				body := util.TirmString(*http.Result.Request.Body)
				return &body
			}
			return nil
		}(),
		Response: func() *string {
			response := util.ToJson(http.Result, false)
			return &response
		}(),
		UserId:  http.UserId,
		Elapsed: http.Elapsed,
	})
}
