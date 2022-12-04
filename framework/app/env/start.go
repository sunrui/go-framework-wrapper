/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-12-04 09:17:22
 */

package env

import "time"

var startAt = time.Now()

// GetStartAt 获取启用时间
func GetStartAt() time.Time {
	return startAt
}
