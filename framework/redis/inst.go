/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-07 00:26:26
 */

package redis

import (
	"config"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

// Inst 实例
var Inst *Redis

func init() {
	// 建立连接池
	Inst = &Redis{
		Pool: redis.Pool{
			MaxIdle:     5,
			MaxActive:   100,
			IdleTimeout: 1 * time.Hour,
			Wait:        true,
			Dial: func() (redis.Conn, error) {
				address := fmt.Sprintf("%s:%d", config.Inst().Redis.Host, config.Inst().Redis.Port)
				timeout := time.Duration(10) * time.Second

				return redis.Dial("tcp", address,
					redis.DialPassword(config.Inst().Redis.Password),
					redis.DialDatabase(config.Inst().Redis.Database),
					redis.DialConnectTimeout(timeout),
					redis.DialReadTimeout(timeout),
					redis.DialWriteTimeout(timeout))
			},
		},
	}

	// 尝试数据库连接
	if _, err := Inst.Pool.Get().Do("PING"); err != nil {
		panic(err.Error())
	}
}
