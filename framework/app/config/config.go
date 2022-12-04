/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-07 00:15:14
 */

package config

import (
	"encoding/json"
	env2 "framework/app/env"
	"framework/app/glog/log"
	"framework/app/mysql"
	"framework/app/redis"
	"framework/app/token"
	"framework/server"
	"os"
)

// Config 配置
type Config struct {
	Mysql  mysql.Config  `json:"mysql"`  // Mysql
	Redis  redis.Config  `json:"redis"`  // Redis
	Server server.Config `json:"server"` // Server
	Log    log.Config    `json:"log"`    // Log
	Token  token.Config  `json:"token"`  // Token
}

// New 创建
func New(jsonFile string) (*Config, error) {
	type env struct {
		Dev  Config `json:"dev"`
		Prod Config `json:"prod"`
	}

	var e env
	if stream, err := os.ReadFile(jsonFile); err != nil {
		return nil, err
	} else if err = json.Unmarshal(stream, &e); err != nil {
		return nil, err
	}

	if env2.IsDev() {
		return &e.Dev, nil
	} else {
		return &e.Prod, nil
	}
}
