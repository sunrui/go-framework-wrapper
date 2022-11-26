/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-26 17:50:58
 */

package mysql

import (
	"fmt"
	"framework/config"
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
	if config.IsDev() {
		fmt.Print(str)
	}
}

// 获取日志
func getLogger(log *logrus.Logger) logger.Interface {
	// 慢日志打印
	var slowThreshold time.Duration

	if config.IsDev() {
		slowThreshold = time.Millisecond // 1 毫秒
	} else {
		slowThreshold = 100 * time.Millisecond // 100 毫秒
	}

	return logger.New(
		&myLog{
			log: log,
		},
		logger.Config{
			SlowThreshold: slowThreshold,
			LogLevel:      logger.Info,
		},
	)
}
