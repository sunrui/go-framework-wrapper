/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-07 00:26:26
 */

package redis

import "medium/env"

var Inst *Redis

func init() {
	Inst = newRedis(env.LoadConfig[config]())
}
