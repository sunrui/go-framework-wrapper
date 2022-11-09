/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-07 00:16:34
 */

package configer

import (
	"flag"
	"testing"
)

// 当前环境
var build *string

// IsDev 是否为开发环境
func IsDev() bool {
	return build != nil && *build != "prod"
}

// IsProd 是否为生产环境
func IsProd() bool {
	return !IsDev()
}

// 初始化
func init() {
	testing.Init()

	// 解析参数，如 -build prod
	flag.Parse()
	build = flag.String("build", "dev", "编译类型")
}
