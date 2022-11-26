/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-16 23:07:01
 */

package log

import (
	"framework/config"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

// HttpAccess 日志实例
var HttpAccess *logrus.Logger

// HttpError 日志实例
var HttpError *logrus.Logger

// Mysql 日志实例
var Mysql *logrus.Logger

func init() {
	// http 访问
	if config.Inst().Log.Switch.HttpAccess {
		HttpAccess = newLog("http", "access")
		// 开启 gin 日志
		gin.DefaultWriter = io.MultiWriter(HttpAccess.Out, os.Stdout)
	}

	// http 错误
	if config.Inst().Log.Switch.HttpError {
		HttpError = newLog("http", "error")
	}

	// mysql
	if config.Inst().Log.Switch.Mysql {
		Mysql = newLog("mysql", "mysql")
	}
}
