/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/09 14:18:09
 */

package config

import (
	"flag"
	"testing"
)

// 当前环境
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
	// 解析参数，如 -build product
	testing.Init()
	flag.Parse()
	build = flag.String("build", "dev", "编译类型")
}
