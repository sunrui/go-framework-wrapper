/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-12-04 17:23:41
 */

package glog

import (
	"fmt"
	"time"
)

// Layout 布局
type Layout interface {
	// 获取布局
	getLayout(level Level, message string) string
}

// DefaultLayout 默认布局
type DefaultLayout struct {
}

// 获取布局
func (defaultLayout DefaultLayout) getLayout(level Level, message string) string {
	timeNow := time.Now().Format("2006-01-02 15:04:05")
	return fmt.Sprintf("%s - %-5s - %s", timeNow, level.String(), message)
}
