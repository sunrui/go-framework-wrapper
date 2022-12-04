/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-26 17:50:58
 */

package mysql

import (
	"framework/app/env"
	"framework/app/glog"
	"gorm.io/gorm/logger"
	"time"
)

// 日志
type myLog struct {
	gLog *glog.GLog // log
}

// Printf 序列化
func (log myLog) Printf(format string, v ...interface{}) {
	log.gLog.Print(glog.DebugLevel, format, v...)
}

// 获取日志
func getLogger(gLog *glog.GLog) logger.Interface {
	return logger.New(
		&myLog{
			gLog: gLog,
		},
		logger.Config{
			SlowThreshold: 50 * time.Millisecond, // 慢查询仅在 warn 级别时才会生效，默认 info 级别下全部输出
			LogLevel: func() logger.LogLevel {
				if env.IsDev() {
					return logger.Info
				} else {
					return logger.Warn
				}
			}(),
		},
	)
}
