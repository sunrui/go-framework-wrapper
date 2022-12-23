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

// New 创建
func New(jsonFile string) (ctx *Context, err error) {
	ctx = &Context{}

	// 初始化配置文件
	if err = ctx.initConfig(jsonFile); err != nil {
		return nil, err
	}

	// 初始化聚合
	type initFunc func() error
	for _, init := range []initFunc{
		ctx.initLogHttpAccess,
		ctx.initLogHttpError,
		ctx.initLogMysql,
		ctx.initService,
		ctx.initMysql,
		ctx.initRedis,
		ctx.initJwtToken,
		ctx.initRedisToken,
		ctx.initRedisStorage,
		ctx.initGinMode,
		ctx.initGinDefaultLog,
		ctx.initGinErrorLog,
	} {
		if err = init(); err != nil {
			return nil, err
		}
	}

	return
}

// 初始化配置文件
func (ctx *Context) initConfig(jsonFile string) (err error) {
	if ctx.Config, err = config.New(jsonFile); err != nil {
		return err
	}

	return nil
}

// 初始化 http 访问日志
func (ctx *Context) initLogHttpAccess() (err error) {
	ctx.Log.HttpAccess = glog.NewGLog(glog.DefaultLayout{}, []glog.Appender{})

	if build.IsDev() {
		ctx.Log.HttpAccess.Appenders = append(ctx.Log.HttpAccess.Appenders, &glog.ConsoleAppender{})
	}

	if ctx.Config.Log.Enable {
		var httpDebugFileLog *log.Log
		if httpDebugFileLog, err = log.New(ctx.Config.Log, "http", "debug"); err != nil {
			return err
		}

		ctx.Log.HttpAccess.Appenders = append(ctx.Log.HttpAccess.Appenders, &glog.FileAppender{
			Debug: httpDebugFileLog,
		})
	}

	return nil
}

// 初始化 http 错误日志
func (ctx *Context) initLogHttpError() (err error) {
	ctx.Log.HttpError = glog.NewGLog(glog.DefaultLayout{}, []glog.Appender{})

	if build.IsDev() {
		ctx.Log.HttpError.Appenders = append(ctx.Log.HttpError.Appenders, &glog.ConsoleAppender{})
	}

	if ctx.Config.Log.Enable {
		var httpErrorFileLog *log.Log
		if httpErrorFileLog, err = log.New(ctx.Config.Log, "http", "error"); err != nil {
			return err
		}

		ctx.Log.HttpError.Appenders = append(ctx.Log.HttpError.Appenders, &glog.FileAppender{
			Error: httpErrorFileLog,
		})
	}

	return nil
}

// 初始化 mysql 访问日志
func (ctx *Context) initLogMysql() (err error) {
	ctx.Log.Mysql = glog.NewGLog(glog.DefaultLayout{}, []glog.Appender{})

	if build.IsDev() {
		ctx.Log.Mysql.Appenders = append(ctx.Log.Mysql.Appenders, &glog.ConsoleAppender{})
	}

	if ctx.Config.Log.Enable {
		var mysqlFileLog *log.Log
		if mysqlFileLog, err = log.New(ctx.Config.Log, "mysql", "mysql"); err != nil {
			return err
		}

		ctx.Log.Mysql.Appenders = append(ctx.Log.Mysql.Appenders, &glog.FileAppender{
			Debug: mysqlFileLog,
		})
	}

	return nil
}

// 初始化 service 日志
func (ctx *Context) initService() (err error) {
	ctx.Log.Service = glog.NewGLog(glog.DefaultLayout{}, []glog.Appender{})

	if build.IsDev() {
		ctx.Log.Service.Appenders = append(ctx.Log.Service.Appenders, &glog.ConsoleAppender{})
	}

	if ctx.Config.Log.Enable {
		var serviceFileLog *log.Log
		if serviceFileLog, err = log.New(ctx.Config.Log, "service", "service"); err != nil {
			return err
		}

		ctx.Log.Service.Appenders = append(ctx.Log.Service.Appenders, &glog.FileAppender{
			Debug: serviceFileLog,
			Info:  serviceFileLog,
			Warn:  serviceFileLog,
			Error: serviceFileLog,
		})
	}

	return nil
}

// 初始化 mysql 数据库
func (ctx *Context) initMysql() (err error) {
	if ctx.Mysql, err = mysql.New(ctx.Config.Mysql, ctx.Log.Mysql); err != nil {
		return err
	}

	return nil
}

// 初始化 redis 缓存
func (ctx *Context) initRedis() (err error) {
	if ctx.Redis, err = redis.New(ctx.Config.Redis); err != nil {
		return err
	}

	return nil
}

// 初始化 jwt 令牌
func (ctx *Context) initJwtToken() (err error) {
	jwtStorage := token.NewJwtStorage([]byte(ctx.Config.Token.JwtSecret))
	ctx.Token.Jwt = token.New(ctx.Config.Token, jwtStorage)

	return nil
}

// 初始化 redis 令牌
func (ctx *Context) initRedisToken() (err error) {
	redisStorage := token.NewRedisStorage(ctx.Redis, ctx.Config.Token.Key)
	ctx.Token.Redis = token.New(ctx.Config.Token, redisStorage)

	return nil
}

// 初始化 redis 令牌
func (ctx *Context) initRedisStorage() (err error) {
	redisStorage := token.NewRedisStorage(ctx.Redis, ctx.Config.Token.Key)
	ctx.Token.Redis = token.New(ctx.Config.Token, redisStorage)

	return nil
}

// gin 环境
func (ctx *Context) initGinMode() (err error) {
	if build.IsDev() {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	return nil
}

// gin 默认日志
func (ctx *Context) initGinDefaultLog() (err error) {
	ginDefaultLog := glog.NewGLog(glog.DefaultLayout{}, []glog.Appender{})

	if build.IsDev() {
		ginDefaultLog.Appenders = append(ginDefaultLog.Appenders, &glog.ConsoleAppender{})
	}

	if ctx.Config.Log.Enable {
		var ginDefaultFileLog *log.Log
		if ginDefaultFileLog, err = log.New(ctx.Config.Log, "gin", "default"); err != nil {
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
func (ctx *Context) initGinErrorLog() (err error) {
	ginDefaultErrorLog := glog.NewGLog(glog.DefaultLayout{}, []glog.Appender{})

	if build.IsDev() {
		ginDefaultErrorLog.Appenders = append(ginDefaultErrorLog.Appenders, &glog.ConsoleAppender{})
	}

	if ctx.Config.Log.Enable {
		var ginErrorFileLog *log.Log
		if ginErrorFileLog, err = log.New(ctx.Config.Log, "gin", "error"); err != nil {
			return err
		}

		ginDefaultErrorLog.Appenders = append(ginDefaultErrorLog.Appenders, &glog.FileAppender{
			Debug: ginErrorFileLog,
		})
	}

	gin.DefaultErrorWriter = io.MultiWriter(ginDefaultErrorLog)

	return nil
}
