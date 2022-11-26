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

// Http 日志实例
var Http *logrus.Logger

// Mysql 日志实例
var Mysql *logrus.Logger

func init() {
	if !config.Inst().Log.Enable {
		return
	}

	Http = newLog("http", "http")
	// 开启 gin 日志
	gin.DefaultWriter = io.MultiWriter(Http.Out, os.Stdout)

	Mysql = newLog("mysql", "mysql")
}
