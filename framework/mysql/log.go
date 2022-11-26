/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-26 17:50:58
 */

package mysql

import (
	"fmt"
	"framework/app/log"
	"gorm.io/gorm/logger"
	"time"
)

// 日志
type myLog struct {
}

func (m *myLog) Printf(format string, v ...interface{}) {
	str := fmt.Sprintf(format, v...)
	log.Inst.Print(str)
	fmt.Print(str)
}

func getLogger() logger.Interface {
	return logger.New(
		&myLog{},
		logger.Config{
			SlowThreshold: time.Millisecond, // 慢日志打印
			LogLevel:      logger.Info,
		},
	)
}
