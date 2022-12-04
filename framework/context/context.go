/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-26 18:53:31
 */

package context

import (
	"framework/app/config"
	"framework/app/log"
	"framework/app/mysql"
	"framework/app/redis"
	"framework/app/token"
)

// Context 上下文
type Context struct {
	Config        *config.Config // 配置
	HttpAccessLog *log.Log       // api 访问日志
	HttpErrorLog  *log.Log       // api 错误日志
	MysqlLog      *log.Log       // mysql 访问日志
	Mysql         *mysql.Mysql   // 数据库
	Redis         *redis.Redis   // 缓存
	JwtToken      *token.Token   // jwt 令牌
	RedisToken    *token.Token   // redis 令牌
}

// New 创建
func New(jsonFile string) (*Context, error) {
	var context = Context{}
	var err error

	// 加载配置文件
	if context.Config, err = config.New(jsonFile); err != nil {
		return nil, err
	}

	// 初始化 api 访问日志
	if context.HttpAccessLog, err = log.New(context.Config.Log, "http", "access"); err != nil {
		return nil, err
	}

	// 初始化 api 错误日志
	if context.HttpErrorLog, err = log.New(context.Config.Log, "http", "error"); err != nil {
		return nil, err
	}

	// 初始化 mysql 访问日志
	if context.MysqlLog, err = log.New(context.Config.Log, "mysql", "mysql"); err != nil {
		return nil, err
	}

	// 初始化 mysql 数据库
	if context.Mysql, err = mysql.New(context.Config.Mysql, context.MysqlLog); err != nil {
		return nil, err
	}

	// 初始化 redis 缓存
	if context.Redis, err = redis.New(context.Config.Redis); err != nil {
		return nil, err
	}

	// 初始化 jwt 令牌
	jwtStorage := token.NewJwtStorage([]byte(context.Config.Token.JwtSecret))
	context.JwtToken = token.New(context.Config.Token, jwtStorage)

	// 初始化 redis 令牌
	redisStorage := token.NewRedisStorage(context.Redis, context.Config.Token.Key)
	context.RedisToken = token.New(context.Config.Token, redisStorage)

	return &context, nil
}
