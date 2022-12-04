/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-12-04 20:01:01
 */

package build

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

// 初始化
func init() {
	testing.Init()

	// 解析参数，如 -build prod
	flag.Parse()
	build = flag.String("build", "dev", "编译类型")
}
