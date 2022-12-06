/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-26 18:53:31
 */

package context

import (
	"framework/app/build"
	"framework/app/config"
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

// 初始化配置文件
func (context *Context) initConfig(jsonFile string) (err error) {
	if context.Config, err = config.New(jsonFile); err != nil {
		return err
	}

	return nil
}

// 初始化 http 访问日志
func (context *Context) initLogHttpAccess() (err error) {
	context.Log.HttpAccess = glog.NewGLog(glog.DefaultLayout{}, []glog.Appender{})

	if build.IsDev() {
		context.Log.HttpAccess.Appenders = append(context.Log.HttpAccess.Appenders, &glog.ConsoleAppender{})
	}

	if context.Config.Log.Enable {
		var httpDebugFileLog *log.Log
		if httpDebugFileLog, err = log.New(context.Config.Log, "http", "debug"); err != nil {
			return err
		}

		context.Log.HttpAccess.Appenders = append(context.Log.HttpAccess.Appenders, &glog.FileAppender{
			Debug: httpDebugFileLog,
		})
	}

	return nil
}

// 初始化 http 错误日志
func (context *Context) initLogHttpError() (err error) {
	context.Log.HttpError = glog.NewGLog(glog.DefaultLayout{}, []glog.Appender{})

	if build.IsDev() {
		context.Log.HttpError.Appenders = append(context.Log.HttpError.Appenders, &glog.ConsoleAppender{})
	}

	if context.Config.Log.Enable {
		var httpErrorFileLog *log.Log
		if httpErrorFileLog, err = log.New(context.Config.Log, "http", "error"); err != nil {
			return err
		}

		context.Log.HttpError.Appenders = append(context.Log.HttpError.Appenders, &glog.FileAppender{
			Error: httpErrorFileLog,
		})
	}

	return nil
}

// 初始化 mysql 访问日志
func (context *Context) initLogMysql() (err error) {
	context.Log.Mysql = glog.NewGLog(glog.DefaultLayout{}, []glog.Appender{})

	if build.IsDev() {
		context.Log.Mysql.Appenders = append(context.Log.Mysql.Appenders, &glog.ConsoleAppender{})
	}

	if context.Config.Log.Enable {
		var mysqlFileLog *log.Log
		if mysqlFileLog, err = log.New(context.Config.Log, "mysql", "mysql"); err != nil {
			return err
		}

		context.Log.Mysql.Appenders = append(context.Log.Mysql.Appenders, &glog.FileAppender{
			Debug: mysqlFileLog,
		})
	}

	return nil
}

// 初始化 service 日志
func (context *Context) initService() (err error) {
	context.Log.Service = glog.NewGLog(glog.DefaultLayout{}, []glog.Appender{})

	if build.IsDev() {
		context.Log.Service.Appenders = append(context.Log.Service.Appenders, &glog.ConsoleAppender{})
	}

	if context.Config.Log.Enable {
		var serviceFileLog *log.Log
		if serviceFileLog, err = log.New(context.Config.Log, "service", "service"); err != nil {
			return err
		}

		context.Log.Service.Appenders = append(context.Log.Service.Appenders, &glog.FileAppender{
			Debug: serviceFileLog,
			Info:  serviceFileLog,
			Warn:  serviceFileLog,
			Error: serviceFileLog,
		})
	}

	return nil
}

// 初始化 mysql 数据库
func (context *Context) initMysql() (err error) {
	if context.Mysql, err = mysql.New(context.Config.Mysql, context.Log.Mysql); err != nil {
		return err
	}

	return nil
}

// 初始化 redis 缓存
func (context *Context) initRedis() (err error) {
	if context.Redis, err = redis.New(context.Config.Redis); err != nil {
		return err
	}

	return nil
}

// 初始化 jwt 令牌
func (context *Context) initJwtToken() (err error) {
	jwtStorage := token.NewJwtStorage([]byte(context.Config.Token.JwtSecret))
	context.Token.Jwt = token.New(context.Config.Token, jwtStorage)

	return nil
}

// 初始化 redis 令牌
func (context *Context) initRedisToken() (err error) {
	redisStorage := token.NewRedisStorage(context.Redis, context.Config.Token.Key)
	context.Token.Redis = token.New(context.Config.Token, redisStorage)

	return nil
}

// 初始化 redis 令牌
func (context *Context) initRedisStorage() (err error) {
	redisStorage := token.NewRedisStorage(context.Redis, context.Config.Token.Key)
	context.Token.Redis = token.New(context.Config.Token, redisStorage)

	return nil
}

// gin 环境
func (context *Context) initGinMode() (err error) {
	if build.IsDev() {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	return nil
}

// gin 默认日志
func (context *Context) initGinDefaultLog() (err error) {
	ginDefaultLog := glog.NewGLog(glog.DefaultLayout{}, []glog.Appender{})

	if build.IsDev() {
		ginDefaultLog.Appenders = append(ginDefaultLog.Appenders, &glog.ConsoleAppender{})
	}

	if context.Config.Log.Enable {
		var ginDefaultFileLog *log.Log
		if ginDefaultFileLog, err = log.New(context.Config.Log, "gin", "default"); err != nil {
			return err
		}

		ginDefaultLog.Appenders = append(ginDefaultLog.Appenders, &glog.FileAppender{
			Debug: ginDefaultFileLog,
		})
	}

	gin.DefaultWriter = io.MultiWriter(ginDefaultLog)

	return nil
}

// gin 默认错误日志
func (context *Context) initGinErrorLog() (err error) {
	ginDefaultErrorLog := glog.NewGLog(glog.DefaultLayout{}, []glog.Appender{})

	if build.IsDev() {
		ginDefaultErrorLog.Appenders = append(ginDefaultErrorLog.Appenders, &glog.ConsoleAppender{})
	}

	if context.Config.Log.Enable {
		var ginErrorFileLog *log.Log
		if ginErrorFileLog, err = log.New(context.Config.Log, "gin", "error"); err != nil {
			return err
		}
		ginDefaultErrorLog.Appenders = append(ginDefaultErrorLog.Appenders, &glog.FileAppender{
			Debug: ginErrorFileLog,
		})
	}

	gin.DefaultErrorWriter = io.MultiWriter(ginDefaultErrorLog)

	return nil
}

// New 创建
func New(jsonFile string) (*Context, error) {
	var context = Context{}
	var err error

	// 初始化配置文件
	if err = context.initConfig(jsonFile); err != nil {
		return nil, err
	}

	// 初始化其它
	type initFunc func() error
	for _, init := range []initFunc{
		context.initLogHttpAccess,
		context.initLogHttpError,
		context.initLogMysql,
		context.initService,
		context.initMysql,
		context.initRedis,
		context.initJwtToken,
		context.initRedisToken,
		context.initRedisStorage,
		context.initGinMode,
		context.initGinDefaultLog,
		context.initGinErrorLog,
	} {
		if err = init(); err != nil {
			return nil, err
		}
	}

	return &context, nil
}
