/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-26 17:50:58
 */

package mysql

import (
	"fmt"
	"framework/app/env"
	"framework/app/log"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm/logger"
	"time"
)

// 日志
type myLog struct {
	log *logrus.Logger
}

// Printf 序列化
func (m myLog) Printf(format string, v ...interface{}) {
	str := fmt.Sprintf(format, v...)

	// 写入日志
	if m.log != nil {
		m.log.Print(str)
	}

	// 写入控制台
	if env.IsDev() {
		fmt.Print(str)
	}
}

// 获取日志
func getLogger(log *log.Log) logger.Interface {
	var logLevel logger.LogLevel
	if env.IsDev() {
		logLevel = logger.Info
	} else {
		logLevel = logger.Warn
	}

	return logger.New(
		&myLog{
			log: log.Logger,
		},
		logger.Config{
			SlowThreshold: 50 * time.Millisecond,
			LogLevel:      logLevel,
		},
	)
}
