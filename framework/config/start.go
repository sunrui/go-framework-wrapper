/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-14 23:13:36
 */

package config

import "time"

var startTime = time.Now()

// GetStartTime 启用时间
func GetStartTime() time.Time {
	return startTime
}
