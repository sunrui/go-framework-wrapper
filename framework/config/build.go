/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022/08/08 19:13:08
 */

package config

import "flag"

var build *string

// IsDev 是否为开发环境
func IsDev() bool {
	return build != nil && *build != "product"
}

// IsProduct 是否为生产环境
func IsProduct() bool {
	return !IsDev()
}

// 初始化
func init() {
	// 解析参数
	flag.Parse()

	// 解析编译参数如 -build product
	build = flag.String("build", "", "编译类型")
}
