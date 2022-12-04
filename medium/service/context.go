/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-12-04 10:00:05
 */

package service

import (
	"framework/context"
	"medium/service/log"
	"path/filepath"
	"runtime"
)

var Ctx *context.Context // Ctx 上下文

// 初始化
func init() {
	var err error

	_, file, _, _ := runtime.Caller(0)
	path := filepath.Dir(file)

	if Ctx, err = context.New(path + "/config.json"); err != nil {
		panic(err.Error())
	}

	// 为 Ctx 加入日志记录
	appender := log.NewAppender(Ctx.Mysql)
	Ctx.Log.HttpAccess.Appenders = append(Ctx.Log.HttpAccess.Appenders, appender)
	Ctx.Log.HttpError.Appenders = append(Ctx.Log.HttpError.Appenders, appender)
}
