/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-12-04 10:00:05
 */

package service

import (
	"framework/context"
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
func (service *Context) initConfig(jsonFile string) (err error) {
	_, file, _, _ := runtime.Caller(0)
	path := filepath.Dir(file)

	service.Context, err = context.New(path + "/" + jsonFile)

	return
}

// 初始化日志附加者
func (service *Context) initLogAppender() {
	appender := log.NewAppender(service.Mysql)
	service.Log.HttpAccess.Appenders = append(service.Log.HttpAccess.Appenders, appender)
	service.Log.HttpError.Appenders = append(service.Log.HttpError.Appenders, appender)
}

// 初始化数据库
func (service *Context) initMirage() {
	service.Mysql.AutoMigrate(
		&log.Http{},
		&user.User{},
		&user.Info{},
		&user.Device{},
		&user.Role{})
}

// NewContext 创建上下文
func NewContext(jsonFile string) (service *Context, err error) {
	var ctx *context.Context

	service = &Context{
		ctx,
	}

	if err = service.initConfig(jsonFile); err != nil {
		return nil, err
	}

	service.initLogAppender()
	service.initMirage()

	return service, nil
}
