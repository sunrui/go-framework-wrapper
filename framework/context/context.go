/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-26 18:53:31
 */

package context

import (
	"framework/app/config"
	"framework/app/env"
	"framework/app/glog"
	"framework/app/glog/log"
	"framework/app/mysql"
	"framework/app/redis"
	"framework/app/token"
	"github.com/gin-gonic/gin"
	"io"
)

// Log 日志
type Log struct {
	HttpAccess *glog.GLog // http 访问日志
	HttpError  *glog.GLog // http 错误日志
	Mysql      *glog.GLog // mysql 日志
	Service    *glog.GLog // service 日志
}

// Token 令牌
type Token struct {
	Jwt   *token.Token // jwt 令牌
	Redis *token.Token // redis 令牌
}

// Context 上下文
type Context struct {
	Config *config.Config // 配置
	Log    Log            // 日志
	Mysql  *mysql.Mysql   // 数据库
	Redis  *redis.Redis   // 缓存
	Token  Token          // 令牌
}

// New 创建
func New(jsonFile string) (*Context, error) {
	var context = Context{}
	var err error

	// 加载 json 配置文件
	if context.Config, err = config.New(jsonFile); err != nil {
		return nil, err
	}

	// 初始化 http 访问日志
	var httpDebugFileLog *log.Log
	if httpDebugFileLog, err = log.New(context.Config.Log, "http", "debug"); err != nil {
		return nil, err
	}

	context.Log.HttpAccess = &glog.GLog{
		Layout: glog.DefaultLayout{},
		Appenders: []glog.Appender{
			&glog.FileAppender{
				Debug: httpDebugFileLog,
			},
		},
	}

	if env.IsDev() {
		context.Log.HttpAccess.Appenders = append(context.Log.HttpAccess.Appenders, &glog.ConsoleAppender{})
	}

	// 初始化 http 错误日志
	var httpErrorFileLog *log.Log
	if httpErrorFileLog, err = log.New(context.Config.Log, "http", "error"); err != nil {
		return nil, err
	}

	context.Log.HttpError = &glog.GLog{
		Layout: glog.DefaultLayout{},
		Appenders: []glog.Appender{
			&glog.FileAppender{
				Error: httpErrorFileLog,
			},
		},
	}

	if env.IsDev() {
		context.Log.HttpError.Appenders = append(context.Log.HttpError.Appenders, &glog.ConsoleAppender{})
	}

	// 初始化 mysql 访问日志
	var mysqlFileLog *log.Log
	if mysqlFileLog, err = log.New(context.Config.Log, "mysql", "mysql"); err != nil {
		return nil, err
	}

	context.Log.Mysql = &glog.GLog{
		Layout: glog.DefaultLayout{},
		Appenders: []glog.Appender{
			&glog.FileAppender{
				Debug: mysqlFileLog,
			},
		},
	}

	if env.IsDev() {
		context.Log.Mysql.Appenders = append(context.Log.Mysql.Appenders, &glog.ConsoleAppender{})
	}

	// 初始化 service 日志
	var serviceFileLog *log.Log
	if serviceFileLog, err = log.New(context.Config.Log, "service", "service"); err != nil {
		return nil, err
	}

	context.Log.Service = &glog.GLog{
		Layout: glog.DefaultLayout{},
		Appenders: []glog.Appender{
			&glog.FileAppender{
				Debug: serviceFileLog,
				Info:  serviceFileLog,
				Warn:  serviceFileLog,
				Error: serviceFileLog,
			},
		},
	}

	if env.IsDev() {
		context.Log.Service.Appenders = append(context.Log.Service.Appenders, &glog.ConsoleAppender{})
	}

	// 初始化 mysql 数据库
	if context.Mysql, err = mysql.New(context.Config.Mysql, context.Log.Mysql); err != nil {
		return nil, err
	}

	// 初始化 redis 缓存
	if context.Redis, err = redis.New(context.Config.Redis); err != nil {
		return nil, err
	}

	// 初始化 jwt 令牌
	jwtStorage := token.NewJwtStorage([]byte(context.Config.Token.JwtSecret))
	context.Token.Jwt = token.New(context.Config.Token, jwtStorage)

	// 初始化 redis 令牌
	redisStorage := token.NewRedisStorage(context.Redis, context.Config.Token.Key)
	context.Token.Redis = token.New(context.Config.Token, redisStorage)

	// gin 环境
	if env.IsDev() {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// gin 默认日志
	var ginDefaultFileLog *log.Log
	if ginDefaultFileLog, err = log.New(context.Config.Log, "gin", "default"); err != nil {
		return nil, err
	}

	ginDefaultLog := &glog.GLog{
		Layout: glog.DefaultLayout{},
		Appenders: []glog.Appender{
			&glog.FileAppender{
				Debug: ginDefaultFileLog,
			},
		},
	}

	if env.IsDev() {
		ginDefaultLog.Appenders = append(ginDefaultLog.Appenders, &glog.ConsoleAppender{})
	}

	gin.DefaultWriter = io.MultiWriter(ginDefaultLog)

	// gin 默认错误日志
	var ginErrorFileLog *log.Log
	if ginErrorFileLog, err = log.New(context.Config.Log, "gin", "error"); err != nil {
		return nil, err
	}

	ginDefaultErrorLog := &glog.GLog{
		Layout: glog.DefaultLayout{},
		Appenders: []glog.Appender{
			&glog.FileAppender{
				Debug: ginErrorFileLog,
			},
		},
	}

	if env.IsDev() {
		ginDefaultErrorLog.Appenders = append(ginDefaultErrorLog.Appenders, &glog.ConsoleAppender{})
	}

	gin.DefaultErrorWriter = io.MultiWriter(ginDefaultErrorLog)

	return &context, nil
}
