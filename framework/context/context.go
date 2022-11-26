/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-26 18:53:31
 */

package context

import (
	"framework/app/token"
	"framework/config"
	"framework/mysql"
	"framework/redis"
)

var Config *config.Config // 配置文件
var Log *Logs             // 日志
var Mysql *mysql.Mysql    // 数据库
var Redis *redis.Redis    // 缓存
var Token *token.Token    // 令牌

func InitContext(jsonFile string) error {
	var err error

	if Config, err = config.NewConfig(jsonFile); err != nil {
		return err
	}

	Log = NewLog(Config.Log)
	Mysql = mysql.NewMysql(Config.Mysql, Log.Mysql)
	Redis = redis.NewRedis(Config.Redis)
	Token = token.NewToken(Config.Token)

	return nil
}
