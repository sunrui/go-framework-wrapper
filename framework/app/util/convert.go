/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-12-22 16:55:03
 */

package util

import "encoding/json"

// ToJson 转换为 json
func ToJson(v any, indent bool) string {
	var jsonByte []byte
	var err error

	if indent {
		jsonByte, err = json.MarshalIndent(v, "", "\t")
	} else {
		jsonByte, err = json.Marshal(v)
	}

	if err != nil {
		panic(err.Error())
	}

	return string(jsonByte)
}
