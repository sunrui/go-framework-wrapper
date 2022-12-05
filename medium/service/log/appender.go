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
	HttpRepository HttpRepository
	Mysql          *mysql.Mysql
}

// NewAppender 创建附加者
func NewAppender(mysql *mysql.Mysql) Appender {
	return Appender{
		Mysql:          mysql,
		HttpRepository: NewHttpRepository(mysql),
	}
}

// Print 打印
func (appender Appender) Print(_ glog.Level, _ string) {
	panic("cannot invoke this method")
}

// PrintMessage 打印消息
func (appender Appender) PrintMessage(format glog.Format) {
	http := Http{
		Level:  format.Level,
		Ip:     format.Request.Ip,
		Method: format.Request.Method,
		Uri:    format.Request.Uri,
		Body: func() *string {
			if format.Request.Body == nil {
				return nil
			} else {
				body := util.TirmString(*format.Request.Body)
				return &body
			}
		}(),
		Response: func() *string {
			if format.Result == nil {
				return nil
			} else {
				result := util.TirmString(format.Result.String())
				return &result
			}
		}(),
		UserId:  format.UserId,
		Elapsed: format.Elapsed,
	}

	appender.Mysql.Save(&http)
}
