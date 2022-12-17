/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-12-16 20:45:49
 */

package glog

import (
	"encoding/json"
	"framework/app/result"
	"framework/app/util"
)

// Http 协议
type Http struct {
	Result  *result.Result // 结果
	UserId  *string        // 用户 id
	Elapsed int64          // 耗时
}

// String 数据
func (http Http) String() string {
	marshal, _ := json.MarshalIndent(http, "", "\t")
	return string(marshal)
}

// LineString 行数据
func (http Http) LineString() string {
	// method http://host:port?query protocol
	buffer := http.Result.Request.Method + " " + http.Result.Request.Uri

	// 空一行
	buffer += "\n"

	// header
	for key, value := range http.Result.Request.Header {
		buffer += key + ": " + value[0] + "\n"
	}

	// 空一行
	buffer += "\n"

	// body
	if http.Result.Request.Body != nil {
		buffer += util.TirmString(*http.Result.Request.Body) + "\n"
	} else {
		buffer += "<null>\n"
	}

	// 空一行
	buffer += "\n"

	// 结果
	if http.Result != nil {
		buffer += util.TirmString(http.Result.String()) + "\n"
	} else {
		buffer += "<null>\n"
	}

	// 空一行
	buffer += "\n"

	return buffer
}
