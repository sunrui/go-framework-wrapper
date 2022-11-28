/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-26 19:23:42
 */

package context

import (
	"framework/app/log"
	"framework/config"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

// Logs 日志
type Logs struct {
	HttpAccess *logrus.Logger // http 访问
	HttpError  *logrus.Logger // http 错误
	Mysql      *logrus.Logger // mysql
}

// NewLogs 创建日志
func NewLogs(logConfig config.Log) *Logs {
	// HttpAccess 日志实例
	var HttpAccessLogger *logrus.Logger

	// HttpError 日志实例
	var HttpErrorLogger *logrus.Logger

	// Mysql 日志实例
	var MysqlLogger *logrus.Logger

	// http 访问
	if logConfig.Switch.HttpAccess {
		HttpAccessLogger = log.New(logConfig, "http", "access")

		// 开启 gin 日志
		if config.IsDev() {
			gin.DefaultWriter = io.MultiWriter(HttpAccessLogger.Out, os.Stdout)
		} else {
			gin.DefaultWriter = io.MultiWriter(HttpAccessLogger.Out)
		}
	}

	// http 错误
	if logConfig.Switch.HttpError {
		HttpErrorLogger = log.New(logConfig, "http", "error")
	}

	// mysql
	if logConfig.Switch.Mysql {
		MysqlLogger = log.New(logConfig, "mysql", "mysql")
	}

	return &Logs{
		HttpAccess: HttpAccessLogger,
		HttpError:  HttpErrorLogger,
		Mysql:      MysqlLogger,
	}
}
