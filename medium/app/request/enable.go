/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-15 00:07:19
 */

package request

// 开关
var enable = true

// SetEnable 设置  开关
func SetEnable(_enable bool) {
	enable = _enable
}

// IsEnable 获取  开关
func IsEnable() bool {
	return enable
}
