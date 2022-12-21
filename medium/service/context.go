/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-12-04 10:00:05
 */

package service

import (
	"framework/context"
	"medium/service/channel"
	"medium/service/log"
	"medium/service/user"
	"path/filepath"
	"runtime"
)

// Context 上下文
type Context struct {
	*context.Context // 上下文
}

// 初始化配置
func (ctx *Context) initConfig() (err error) {
	_, file, _, _ := runtime.Caller(0)
	path := filepath.Dir(file)

	ctx.Context, err = context.New(path + "/context.json")

	return
}

// 初始化日志附加者
func (ctx *Context) initLogAppender() {
	appender := log.NewAppender(ctx.Mysql)
	ctx.Log.HttpAccess.Appenders = append(ctx.Log.HttpAccess.Appenders, appender)
	ctx.Log.HttpError.Appenders = append(ctx.Log.HttpError.Appenders, appender)
}

// 初始化数据库
func (ctx *Context) initMirage() {
	ctx.Mysql.AutoMigrate(
		&channel.Channel{},
		&log.LogHttp{},
		&user.User{},
		&user.UserInfo{},
		&user.UserDevice{},
		&user.UserRole{})
}

// NewContext 创建上下文
func NewContext() (ctx *Context, err error) {
	ctx = &Context{}

	if err = ctx.initConfig(); err != nil {
		return nil, err
	}

	ctx.initLogAppender()
	ctx.initMirage()

	return ctx, nil
}
