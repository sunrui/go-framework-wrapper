/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-12-04 20:01:02
 */

package build

import "time"

// 启动时间
var startAt = time.Now()

// GetStartAt 获取启动时间
func GetStartAt() time.Time {
	return startAt
}
