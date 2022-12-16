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
	// 获取消息布局
	getHttpLayout(level Level, http Http) string
}

// DefaultLayout 默认布局
type DefaultLayout struct {
}

// 获取布局
func (defaultLayout DefaultLayout) getLayout(level Level, message string) string {
	timeNow := time.Now().Format("2006-01-02 15:04:05")
	return fmt.Sprintf("%s - %-5s - %s", timeNow, level.String(), message)
}

// 获取消息布局
func (defaultLayout DefaultLayout) getHttpLayout(level Level, http Http) string {
	timeNow := time.Now().Format("2006-01-02 15:04:05")

	var UserId string
	if http.UserId != nil {
		UserId = fmt.Sprintf(" - userId(%s)", *http.UserId)
	} else {
		UserId = ""
	}

	return fmt.Sprintf("%s - %-5s - %dms - %s%s\n%s",
		timeNow, level.String(), http.Elapsed, http.Result.Request.Ip, UserId, http.LineString())
}
