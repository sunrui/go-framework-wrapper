/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-12-04 10:00:05
 */

package service

import (
	"framework/context"
	"path/filepath"
	"runtime"
)

var Ctx *context.Context // Ctx 上下文

// 初始化
func init() {
	var err error

	_, file, _, _ := runtime.Caller(0)
	path := filepath.Dir(file)

	if Ctx, err = context.New(path + "/config.json"); err != nil {
		panic(err.Error())
	}
}
