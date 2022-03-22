/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022/01/01
 */

package util

import (
	"io/ioutil"
	"os"
)

// ReadStream 加载配置文件流
func ReadStream(file string) ([]byte, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(f)
}
