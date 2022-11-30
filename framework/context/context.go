/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-26 18:53:31
 */

package context

import (
	"framework/config"
	"framework/http/token"
	"framework/mysql"
	"framework/redis"
)

var Config *config.Config // 配置
var Log *Logs             // 日志
var Mysql *mysql.Mysql    // 数据库
var Redis *redis.Redis    // 缓存
var Token *token.Token    // 令牌

// Init 初始化
func Init(jsonFile string) error {
	var err error

	// 初始化配置
	if Config, err = config.New(jsonFile); err != nil {
		return err
	}

	// 初始化日志
	Log = NewLogs(Config.Log)

	// 初始化 mysql 数据库
	Mysql = mysql.New(Config.Mysql, Log.Mysql)

	// 初始化 redis 缓存
	Redis = redis.New(Config.Redis)

	// 初始化令牌
	tokenStorage := token.NewRedisStorage(Redis, Config.Token.Key)
	Token = token.New(Config.Token, tokenStorage)

	return nil
}
