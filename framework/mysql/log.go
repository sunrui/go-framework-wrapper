/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-26 17:50:58
 */

package mysql

import (
	"fmt"
	"framework/app/log"
	"framework/config"
	"gorm.io/gorm/logger"
	"time"
)

// 自定义日志
type myLog struct {
}

// Printf 序列化
func (m *myLog) Printf(format string, v ...interface{}) {
	str := fmt.Sprintf(format, v...)

	// 写入日志
	log.Mysql.Print(str)
	// 写入控制台
	fmt.Print(str)
}

// 获取日志对象
func getLogger() logger.Interface {
	// 慢日志打印
	var slowThreshold time.Duration

	if config.IsDev() {
		slowThreshold = time.Millisecond // 1 毫秒
	} else {
		slowThreshold = 100 * time.Millisecond // 100 毫秒
	}

	return logger.New(
		&myLog{},
		logger.Config{
			SlowThreshold: slowThreshold,
			LogLevel:      logger.Info,
		},
	)
}
